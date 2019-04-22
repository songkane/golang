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
