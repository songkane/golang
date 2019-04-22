// Kafka kafka 通用写kafka接口
// Created by chenguolin 2018-09-21
package kafka

import (
	"fmt"
	"testing"
)

func TestNewSyncKafkaProducer(t *testing.T) {
	// case 1
	brokers := ""
	producer, err := NewSyncKafkaProducer(brokers)
	if err == nil {
		t.Fatal("TestNewSyncKafkaProducer case 1 err == nil")
	}
	if producer != nil {
		t.Fatal("TestNewSyncKafkaProducer case 1 producer != nil")
	}

	// case 2
	brokers = "127.0.0.1:9092"
	producer, err = NewSyncKafkaProducer(brokers)
	fmt.Println(err)
	if err != nil {
		t.Fatal("TestNewSyncKafkaProducer case 2 err != nil")
	}
	if producer == nil {
		t.Fatal("TestNewSyncKafkaProducer case 2 producer == nil")
	}
}
