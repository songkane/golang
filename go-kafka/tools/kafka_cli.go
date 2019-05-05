// Package main kafka console
// Created by chenguolin 2019-04-20
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"gitlab.local.com/golang/go-kafka"
	"gitlab.local.com/golang/go-kafka/pkg/sarama"
	cluster "gitlab.local.com/golang/go-kafka/pkg/sarama-cluster"
)

func main() {
	// 1. parse args
	args := parseArgs()
	if args.h || args.help || (!args.help && !args.list && !args.query &&
		!args.consumer && !args.producer) {
		usage()
		return
	}

	// 2. command
	if args.list {
		listCmd(args)
	} else if args.query {
		queryCmd(args)
	} else if args.consumer {
		consumerCmd(args)
	} else if args.producer {
		producerCmd(args)
	} else {
		usage()
	}
}

// CliArgs kafka cli args
type CliArgs struct {
	// cmd
	h        bool //help
	help     bool //help
	list     bool //list topic
	query    bool //query topic
	consumer bool //consumer message
	producer bool //producer message
	// options
	brokers   string //brokers
	topic     string //topics
	group     string //consumer group
	partition int64  //partition
	replica   int64  //replica
	offset    int64  //offset
	key       string //produce message key
	value     string //produce message value
}

func parseArgs() *CliArgs {
	args := &CliArgs{}
	// cmd
	flag.BoolVar(&args.h, "h", false, "help.")
	flag.BoolVar(&args.help, "help", false, "help.")
	flag.BoolVar(&args.list, "list", false, "show topic list.")
	flag.BoolVar(&args.query, "query", false, "query topic info.")
	flag.BoolVar(&args.consumer, "consume", false, "consumer message.")
	flag.BoolVar(&args.producer, "produce", false, "producer message.")

	// options
	flag.StringVar(&args.brokers, "brokers", "", "broker list, like localhost:9092.")
	flag.StringVar(&args.topic, "topic", "", "topic name.")
	flag.StringVar(&args.group, "group", "", "consumer group id.")
	flag.Int64Var(&args.partition, "partition", 0, "partition.")
	flag.Int64Var(&args.replica, "replica", 1, "replica.")
	flag.Int64Var(&args.offset, "offset", 0, "offset.")
	flag.StringVar(&args.key, "key", "", "produce message key.")
	flag.StringVar(&args.value, "value", "", "produce message value.")

	flag.Parse()
	return args
}

func usage() {
	fmt.Println("kafka-cli is a console util tool to access kafka cluster.")
	fmt.Println("")

	fmt.Println("Usage:")
	fmt.Println("\tkafak-cli [command]")
	fmt.Println("")

	fmt.Println("Available Commands:")
	fmt.Println("\t-h/-help\thelp about any command.")
	fmt.Println("\t-list\t\tlist all topics.")
	fmt.Println("\t-query\t\tquery topics info.")
	fmt.Println("\t-consume\tconsume from kafka topic.")
	fmt.Println("\t-produce\tproduce message 2 kafka topic.")
	fmt.Println("")

	fmt.Println("Options:")
	fmt.Println("\t-brokers\tbroker list, like 127.0.0.1:9092,127.0.0.2:9092.")
	fmt.Println("\t-topic\t\ttopic name.")
	fmt.Println("\t-group\t\tconsumer group name.")
	fmt.Println("\t-partition\tpartition count or partition id, default 0.")
	fmt.Println("\t-replica\treplica count, default 1.")
	fmt.Println("\t-offset\t\tconsume offset 0=newest 1=oldest, default 0.")
	fmt.Println("\t-key\t\tproduce message key.")
	fmt.Println("\t-valu3\t\tproduce message value.")
}

// list command
func listCmd(args *CliArgs) {
	if args.brokers == "" {
		fmt.Println("List Topics: kafka-cli -list -bokers localhost:9092")
		fmt.Println("")
		usage()
		return
	}

	// new client
	client, err := kafka.NewClient(args.brokers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// list topic
	topics, err := client.Topics()
	if err != nil {
		fmt.Println(err)
		return
	}

	// print topics
	for idx, topic := range topics {
		fmt.Printf("%d. %s\n", idx+1, topic)
	}
}

// query topic
func queryCmd(args *CliArgs) {
	if args.brokers == "" || args.topic == "" {
		fmt.Println("Query Topic: kafka-cli -query -bokers localhost:9092 -topic xxxx")
		fmt.Println("")
		usage()
		return
	}

	// new client
	client, err := kafka.NewClient(args.brokers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// list topic
	topics, err := client.Topics()
	if err != nil {
		fmt.Println(err)
		return
	}
	// print topics
	for idx, topic := range topics {
		if topic != args.topic {
			continue
		}

		partitions, err := client.Partitions(topic)
		if err != nil {
			fmt.Printf("%d. %s, partitions: %v\n", idx+1, topic, partitions)
			continue
		}

		for _, partition := range partitions {
			// get offset
			oldOffset, _ := client.GetOffset(topic, partition, sarama.OffsetOldest)
			newOffset, _ := client.GetOffset(topic, partition, sarama.OffsetNewest)
			// get replica
			replicas, _ := client.Replicas(topic, partition)

			info := fmt.Sprintf("partition:%d, oldOffset:%d, newOffset:%d, replicas:%v",
				partition, oldOffset, newOffset, replicas)
			fmt.Println(info)
		}
	}
}

// consumer command
func consumerCmd(args *CliArgs) {
	brokers := args.brokers
	topic := args.topic
	groupID := args.group
	partition := args.partition
	offset := args.offset

	if brokers == "" || topic == "" {
		fmt.Println("kafka-cli -c -bokers localhost:9092 -topic xxxx -group ...")
		usage()
		return
	}

	// random generate groupID
	if groupID == "" {
		name, err := os.Hostname()
		if err != nil {
			name = "unknown"
		}
		currentMilliSec := time.Now().UnixNano() / int64(time.Millisecond)
		randBytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			randBytes[i] = byte(rand.Intn(26) + 65)
		}
		groupID = fmt.Sprintf("%s-%d-%s", name, currentMilliSec, string(randBytes))
	}

	// new partition consumer
	cfg := cluster.NewConfig()
	cfg.Consumer.Return.Errors = true
	// Users who require low-level access can enable ConsumerModePartitions where individual partitions
	// are exposed on the Partitions() channel. Messages and errors must then be consumed on the partitions
	// themselves.
	cfg.Group.Mode = cluster.ConsumerModePartitions
	if offset == kafka.OffsetOldest {
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	} else {
		cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	topics := []string{topic}
	consumer, err := cluster.NewConsumer(strings.Split(brokers, ","), groupID, topics, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// goroutine process
	for {
		select {
		case part := <-consumer.Partitions():
			if int64(part.Partition()) == partition {
				// start a separate goroutine to consume messages
				go func(pc cluster.PartitionConsumer) {
					for msg := range pc.Messages() {
						v := fmt.Sprintf("[%s/%d/%d]\t{\"key\":%s,\"message\":%s}", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
						fmt.Println(v)
						// mark message as processed
						consumer.MarkOffset(msg, "")
					}
				}(part)
			}
		case <-signals:
			return
		}
	}
}

// producer command
func producerCmd(args *CliArgs) {
	brokers := args.brokers
	topic := args.topic
	key := args.key
	value := args.value

	if brokers == "" || topic == "" || value == "" {
		fmt.Println("kafka-cli -p -bokers localhost:9092 -topic xxxx -key xxxx -value xxxx")
		usage()
		return
	}

	// new produce
	producer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer producer.Close()

	partition, offset, err := producer.Send(topic, key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	v := fmt.Sprintf("[%s/%d/%d] ok.", topic, partition, offset)
	fmt.Println(v)
}
