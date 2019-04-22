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
