// Package kafka consumer unit test
// Created by chenguolin 2019-04-20
package kafka

import (
	"fmt"
	"testing"
	"time"

	"gitlab.local.com/golang/go-kafka/pkg/sarama"
)

func TestNewConsumer(t *testing.T) {
	// case 1
	topic := ""
	groupID := ""
	brokers := ""
	defaultOffset := OffsetNewset

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err == nil {
		t.Fatal("TestNewConsumer case 1 err != nil")
	}
	if consumer != nil {
		t.Fatal("TestNewConsumer case 1 consumer == nil")
	}

	// case 2
	topic = "k8s-log-test-output-stdout"
	groupID = ""
	brokers = ""
	defaultOffset = OffsetOldest

	consumer, err = NewConsumer(brokers, topic, groupID, defaultOffset)
	if err == nil {
		t.Fatal("TestNewConsumer case 2 err != nil")
	}
	if consumer != nil {
		t.Fatal("TestNewConsumer case 2 consumer == nil")
	}

	// case 3
	topic = "k8s-log-test-output-stdout"
	groupID = "TestNewConsumer"
	brokers = ""
	defaultOffset = OffsetNewset

	consumer, err = NewConsumer(brokers, topic, groupID, defaultOffset)
	if err == nil {
		t.Fatal("TestNewConsumer case 3 err != nil")
	}
	if consumer != nil {
		t.Fatal("TestNewConsumer case 3 consumer == nil")
	}

	// case 4
	topic = "k8s-log-test-output-stdout"
	groupID = "TestNewConsumer"
	brokers = "localhost:9092,localhost:9092"
	defaultOffset = OffsetOldest

	consumer, err = NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestNewConsumer case 4 err != nil")
	}
	if consumer == nil {
		t.Fatal("TestNewConsumer case 4 consumer == nil")
	}

	consumer.Close()
}

func TestConsumer_Messages(t *testing.T) {
	topic := "k8s-log-test-output-stdout"
	groupID := "TestConsumer_Messages"
	brokers := "localhost:9092,localhost:9092"
	defaultOffset := OffsetNewset

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestConsumer_Messages err != nil")
	}
	if consumer == nil {
		t.Fatal("TestConsumer_Messages consumer == nil")
	}
	defer consumer.Close()

	// start consumer
	msgChan := consumer.Messages()
	if msgChan == nil {
		t.Fatal("TestConsumer_Messages msgChan == nil")
	}

	go func(<-chan *sarama.ConsumerMessage) {
		for msg := range msgChan {
			//bytes, _ := json.Marshal(msg)
			//fmt.Println(string(bytes))
			fmt.Println(msg.Offset)
			consumer.CommitOffset(msg)
		}
	}(msgChan)

	time.Sleep(10 * time.Second)
}

func TestConsumer_Errors(t *testing.T) {
	topic := "k8s-log-test-output-stdout"
	groupID := "TestConsumer_Errors"
	brokers := "localhost:9092,localhost:9092"
	defaultOffset := OffsetOldest

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestConsumer_Errors err != nil")
	}
	if consumer == nil {
		t.Fatal("TestConsumer_Errors consumer == nil")
	}
	defer consumer.Close()

	// start consumer
	errChan := consumer.Errors()
	if errChan == nil {
		t.Fatal("TestConsumer_Errors errChan == nil")
	}

	go func(<-chan error) {
		for err := range errChan {
			fmt.Println(err)
		}
	}(errChan)

	time.Sleep(10 * time.Second)
}

func TestConsumer_case1(t *testing.T) {
	topic := "k8s-log-test-output-stdout"
	groupID := "TestConsumer_case1"
	brokers := "localhost:9092,localhost:9092"
	defaultOffset := OffsetOldest

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestConsumer_case1 err != nil")
	}
	if consumer == nil {
		t.Fatal("TestConsumer_case1 consumer == nil")
	}
	defer consumer.Close()

	// consumer
	msgChan := consumer.Messages()
	errChan := consumer.Errors()

	// goroutine
	go func() {
		for {
			select {
			case msg, ok := <-msgChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(string(msg.Topic))
				fmt.Println(msg.Offset)
				fmt.Println(msg.Partition)
				consumer.CommitOffset(msg)
			case err, ok := <-errChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(err.Error())
			}
		}
	}()

	time.Sleep(time.Second * 20)
}

func TestConsumer_case2(t *testing.T) {
	topic := "k8s-log-test-output-stdout"
	groupID := "TestConsumer_case2"
	brokers := "localhost:9092,localhost:9092"
	defaultOffset := OffsetNewset

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestConsumer_case2 err != nil")
	}
	if consumer == nil {
		t.Fatal("TestConsumer_case2 consumer == nil")
	}
	defer consumer.Close()

	// start consumer
	msgChan := consumer.Messages()
	errChan := consumer.Errors()

	// goroutine
	go func() {
		for {
			select {
			case msg, ok := <-msgChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(string(msg.Topic))
				fmt.Println(msg.Offset)
				fmt.Println(msg.Partition)
				consumer.CommitOffset(msg)
			case err, ok := <-errChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(err.Error())
			}
		}
	}()

	time.Sleep(time.Second * 20)
}

func TestConsumer_consumeLoop(t *testing.T) {
	topic := "kafka_topic_test"
	groupID := "TestConsumer_consumeLoop"
	brokers := "localhost:9092"
	defaultOffset := OffsetNewset

	consumer, err := NewConsumer(brokers, topic, groupID, defaultOffset)
	if err != nil {
		t.Fatal("TestConsumer_consumeLoop err != nil")
	}
	if consumer == nil {
		t.Fatal("TestConsumer_consumeLoop consumer == nil")
	}
	defer consumer.Close()

	// start consumer
	msgChan := consumer.Messages()
	errChan := consumer.Errors()

	// goroutine
	go func() {
		for {
			select {
			case msg, ok := <-msgChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(string(msg.Topic), msg.Partition, msg.Offset)
				consumer.CommitOffset(msg)
			case err, ok := <-errChan:
				// channel has closed
				if !ok {
					return
				}
				fmt.Println(err.Error())
			}
		}
	}()

	time.Sleep(1000 * time.Second)
}
