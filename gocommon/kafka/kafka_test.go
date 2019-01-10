// Kafka kafka 通用写kafka接口
// Created by chenguolin 2018-09-21
package kafka

import (
	"testing"
)

func TestNewSyncKafkaProducer(t *testing.T) {
	// case 1
	brokers := []string{}
	producer, err := NewSyncKafkaProducer(brokers)
	if err == nil {
		t.Fatal("TestNewSyncKafkaProducer case 1 err == nil")
	}
	if producer != nil {
		t.Fatal("TestNewSyncKafkaProducer case 1 producer != nil")
	}

	// case 2
	/*
		brokers = []string{"172.16.30.2:9092"}
		producer, err = NewSyncKafkaProducer(brokers)
		fmt.Println(err)
		if err != nil {
			t.Fatal("TestNewSyncKafkaProducer case 2 err != nil")
		}
		if producer == nil {
			t.Fatal("TestNewSyncKafkaProducer case 2 producer == nil")
		}
	*/
}
