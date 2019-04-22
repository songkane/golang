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
