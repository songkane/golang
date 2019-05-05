# go-kafka
golang kafka封装，目前业界用的最多的Go kafka库是https://github.com/Shopify/sarama 但是sarama存在一些问题，所以基于https://github.com/bsm/sarama-cluster 封装了更`傻瓜`
化的kafka调用

1. sarama存在几个问题
    * 对consumer group支持不是很友好，没有提供优化的api调用
    * consumer group没有及时commit offset，导致程序重启的时候会从最新的消息开始消费，导致数据丢失
    * 不支持rebalance机制
2. sarama-cluster实现了上诉提到的几个问题，提供优化的api调用，支持自动commit offset、支持rebalance机制

# kafka简介
1. kafka的队列模型
    * 每个 topic 可以有多个 partition，每个 partition 才是严格意义上的消息队列（消息先进先出），而不同 partition 之间是互不影响的。
      举个例子来说，有消息 A 和消息 B，如果要保证 A 一定要比 B 先被消费，就必须要保证 A 一定要先被投递到某个 partition，且 B 再被投递到同一个 partition。
      如果 B 被投递到了不一样 partition，那么 B 是有可能先于 A 被消费的。
    * “一个 topic 的一个 partition” 是保证无重复消息的最小消费单元，换句话说，如果有两个消费程序消费同一个 topic 的同一个 partition，那它们消费的消息事实上是彼此重复的。
      所以为保证所有 partition 均被消费，且不会被同一个业务（属于同一个 group）的多个消费程序重复消费，是需要一个分配策略来决定每个消费程序应当消费哪几个或哪一个 partition，又或者应当做作为候补（当消费程序数量大于 partition 数量时发生），而 kafka 是希望用户程序自行实现这个分配策略的。

2. kafka应用场景举例
    * 消息没有严格有序性要求，任意消息可以被任意消费程序消费，且消费各消息耗时相近
      可以将消息投递到任意 partition，partition 任意均等分配到各个消费程序；
    * 所有消息要求严格有序，但消息量不大
      可以配置仅一个 partition，一个消费程序负责消费，其他消费程序作为替补；
    * 同组消息要求有序，不同组消息不要求有序，例如同一个用户的消费要求有序，不同用户的消息不要求有序
      可以将属于同一组的消息投递到同一个 partition，比如拿 UID 对 partition 数量取模；
    * 特定组的消息仅可以被特定消费程序消费
      可以在将该特定组的消息投递到特定 partition，配置时指定到特定消费程序；
    * 某些消息消费耗时长且要求有序，有些消息消费耗时短且不要求有序
      可以将分为两组 partition，一组实行针对有序消息的策略，且多一些 partition、多一些消费程序增大处理能力，另一组实行针对无序消息的的策略，且少一些 partition、少一些消费程序节省资源。

# examples
## consumer
```
// Package main kafka consumer example
// Created by chenguoin 2019-04-20
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"

	"gitlab.local.com/golang/go-kafka"
	golog "gitlab.local.com/golang/go-log"
)

func main() {
	fmt.Println("Consumer start ...")

	// new consumer
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	topic := "k8s-log-test-output-stdout"
	groupID := "consumer_example"
	defaultOffset := kafka.OffsetNewset

	// new consumer
	consumer, err := kafka.NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		fmt.Println("kafka.NewConsumer error: ", err.Error())
		os.Exit(1)
	}
	defer consumer.Close()

	// goroutine receive message
	wg := &sync.WaitGroup{}
	wg.Add(1)
	stopChan := make(chan struct{})
	go consume(consumer, stopChan, wg)

	// wait signal
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// blocking unit receive signal
	<-sigchan
	close(stopChan)
	wg.Wait()

	fmt.Println("Consumer stop successfuly ~")
}

func consume(consumer *kafka.Consumer, stopChan chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			golog.Error("consume handle panic",
				golog.Object("error", err))
			debug.PrintStack()
		}
		// goroutine done
		wg.Done()
	}()

	// get message, error channel
	msgChan := consumer.Messages()
	errChan := consumer.Errors()

	for {
		select {
		case msg := <-msgChan:
			bytes, _ := json.Marshal(msg)
			fmt.Println(string(bytes))
			// commit offset 2 zk
			consumer.CommitOffset(msg)
		case err := <-errChan:
			fmt.Println("receive error: ", err.Error())
		case <-stopChan:
			fmt.Println("closing consume ...")
			return
		}
	}
}
```

## sync producer
```
// Package main kafka sync producer example
// Created by chenguoin 2019-04-20
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"gitlab.local.com/golang/go-kafka"
	golog "gitlab.local.com/golang/go-log"
)

func main() {
	fmt.Println("Producer start ...")
	// new sync producer
	brokers := "192.168.0.1:9092,192.168.0.2:9092"

	producer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		fmt.Println("kafka.NewSyncProducer error: ", err.Error())
		os.Exit(1)
	}

	// sync produce message
	go syncProduce(producer)

	// wait signal
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Await a sigterm signal before safely closing the consumer
	<-sigchan

	fmt.Println("Producer stop successfuly ~")
}

func syncProduce(producer *kafka.SyncProducer) {
	defer func() {
		if err := recover(); err != nil {
			golog.Error("syncProduce handle panic",
				golog.Object("error", err))
			debug.PrintStack()
		}
	}()

	topic := "k8s-log-test-output-stdout"
	for {
		// rand key
		randKey := make([]byte, 16)
		for i := 0; i < 16; i++ {
			randKey[i] = byte(rand.Intn(26) + 65)
		}

		// rand value
		randValue := make([]byte, 64)
		for i := 0; i < 64; i++ {
			randValue[i] = byte(rand.Intn(26) + 65)
		}

		partition, offset, err := producer.Send(topic, string(randKey), string(randValue))
		if err != nil {
			golog.Error("producer.Send error", golog.Object("error", err))
			continue
		}

		golog.Info("producer.Send success", golog.Int32("partition", partition),
			golog.Int64("offset", offset))
		time.Sleep(time.Second)
	}
}
```

## async producer
```
// Package main kafka async producer example
// Created by chenguoin 2019-04-20
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"gitlab.local.com/golang/go-kafka"
	golog "gitlab.local.com/golang/go-log"
)

func main() {
	fmt.Println("Producer start ...")
	// new async producer
	brokers := "192.168.0.1:9092,192.168.0.2:9092"

	producer, err := kafka.NewAsyncProducer(brokers)
	if err != nil {
		fmt.Println("kafka.NewAsyncProducer error: ", err.Error())
		os.Exit(1)
	}

	// async produce message
	go asyncProduce(producer)

	// wait signal
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Await a sigterm signal before safely closing the producer
	<-sigchan

	fmt.Println("Producer stop successfuly ~")
}

func asyncProduce(producer *kafka.AsyncProducer) {
	defer func() {
		if err := recover(); err != nil {
			golog.Error("asyncProduce handle panic",
				golog.Object("error", err))
			debug.PrintStack()
		}
	}()

	topic := "k8s-log-test-output-stdout"
	for {
		// rand key
		randKey := make([]byte, 16)
		for i := 0; i < 16; i++ {
			randKey[i] = byte(rand.Intn(26) + 65)
		}

		// rand value
		randValue := make([]byte, 64)
		for i := 0; i < 64; i++ {
			randValue[i] = byte(rand.Intn(26) + 65)
		}

		producer.Send(topic, string(randKey), string(randValue))

		time.Sleep(time.Second)
	}
}
```

## kafka-cli
```
kafka-cli is a console util tool to access kafka cluster.

Usage:
	kafak-cli [command]

Available Commands:
	-h/-help	help about any command.
	-list		list all topics.
	-query		query topics info.
	-consume	consume from kafka topic.
	-produce	produce message 2 kafka topic.

Options:
	-brokers	broker list, like 127.0.0.1:9092,127.0.0.2:9092.
	-topic		topic name.
	-group		consumer group name.
	-partition	partition count or partition id, default 0.
	-replica	replica count, default 1.
	-offset		consume offset 0=newest 1=oldest, default 0.
	-key		produce message key.
	-valu3		produce message value.
```

1. 列出当前所有topics  
   ./kafka_cli -list -brokers 192.168.0.1:9092,192.168.0.2:9092
2. 查询当前topic信息   
   ./kafka_cli -query -brokers 192.168.0.1:9092,192.168.0.2:9092 -topic k8s-log-test-output-stdout
3. 消费某个topic  
   ./kafka_cli -consume -brokers 192.168.0.1:9092,192.168.0.2:9092 -topic k8s-log-test-output-stdout -group console_consumer -partition 0 -offset 2310

   -brokers: 必须字段  
   -topics: 必须字段  
   -group: 可选字段，默认随机生成  
   -partition: 可选字段，默认为0  
   -offset: 可选字段，默认从最新开始读
4. 写数据到某个topic  
   ./kafka_cli -produce -brokers 192.168.0.1:9092,192.168.0.2:9092 -topic k8s-log-test-output-stdout -value "test kafka_cli by cgl"




