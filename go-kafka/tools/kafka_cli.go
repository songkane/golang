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

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"gitlab.local.com/golang/go-kafka"
)

func main() {
	// 1. parse args
	args := parseArgs()
	if args.help || (!args.help && !args.list && !args.consumer && !args.producer) {
		usage()
		return
	}

	// 2. command
	if args.list {
		listCmd(args)
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
	help     bool //help
	list     bool //list
	consumer bool //consumer
	producer bool //producer
	// options
	brokers   string //brokers
	topic     string //consume topics
	group     string //consume group
	partition int64  //consumne partition
	offset    int64  //consume offset
	key       string //produce message key
	value     string //produce message value
}

func parseArgs() *CliArgs {
	args := &CliArgs{}
	// cmd
	flag.BoolVar(&args.help, "h", false, "help.")
	flag.BoolVar(&args.list, "l", false, "show topic list.")
	flag.BoolVar(&args.consumer, "c", false, "consumer message.")
	flag.BoolVar(&args.producer, "p", false, "producer message.")
	// options
	flag.StringVar(&args.brokers, "brokers", "", "broker list, like localhost:9092.")
	flag.StringVar(&args.topic, "topic", "", "consume topic name.")
	flag.StringVar(&args.group, "group", "", "consume group id.")
	flag.Int64Var(&args.partition, "partition", 0, "consume partition id.")
	flag.Int64Var(&args.offset, "offset", 0, "consume offset.")
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
	fmt.Println("\t-h\thelp about any command.")
	fmt.Println("\t-l\tlist all topics.")
	fmt.Println("\t-c\tconsume from kafka topic.")
	fmt.Println("\t-p\tproduce message 2 kafka topic.")

	fmt.Println("Options:")
	fmt.Println("\t-brokers\tbroker list, like 127.0.0.1:9092,127.0.0.2:9092.")
	fmt.Println("\t-topic\t\tconsume topic name.")
	fmt.Println("\t-group\t\tconsume group name.")
	fmt.Println("\t-partition\tconsume partition id, default read from partition 0.")
	fmt.Println("\t-offset\t\tconsume offset 0=newest 1=oldest, default read from newest.")
	fmt.Println("\t-key\t\tproduce message key.")
	fmt.Println("\t-valu3\t\tproduce message value.")
}

// list command
func listCmd(args *CliArgs) {
	if args.brokers == "" {
		fmt.Println("kafka-cli -l -bokers localhost:9092")
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
		partitions, err := client.Partitions(topic)
		if err != nil {
			fmt.Printf("%d. %s, partitions: %v\n", idx+1, topic, partitions)
			continue
		}

		partitionInfo := make([]string, 0)
		for _, partition := range partitions {
			oldOffset, _ := client.GetOffset(topic, partition, sarama.OffsetOldest)
			newOffset, _ := client.GetOffset(topic, partition, sarama.OffsetNewest)
			info := fmt.Sprintf("{partition:%d, oldOffset:%d, newOffset:%d}", partition, oldOffset, newOffset)
			partitionInfo = append(partitionInfo, info)
		}
		fmt.Printf("%d. %s, %v\n", idx+1, topic, partitionInfo)
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
