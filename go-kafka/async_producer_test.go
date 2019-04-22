// Package kafka async producer unit test
// Created by chenguolin 2019-04-20
package kafka

import (
	"testing"
	"time"
)

func TestNewAsyncProducer_case1(t *testing.T) {
	asyncProducer, err := NewAsyncProducer("")
	if err == nil {
		t.Fatal("TestNewAsyncProducer_case1 err == nil")
	}
	if asyncProducer != nil {
		t.Fatal("TestNewAsyncProducer_case1 asyncProducer != nil")
	}
}

func TestNewAsyncProducer_case2(t *testing.T) {
	brokers := "172.16.28.47:9092,172.16.28.48:9092,172.16.28.49:9092"
	asyncProducer, err := NewAsyncProducer(brokers)
	if err != nil {
		t.Fatal("TestNewAsyncProducer_case2 err == nil")
	}
	if asyncProducer == nil {
		t.Fatal("TestNewAsyncProducer_case2 asyncProducer != nil")
	}
	defer asyncProducer.Close()
}

func TestAsyncProducer_Send(t *testing.T) {
	brokers := "172.16.28.47:9092,172.16.28.48:9092,172.16.28.49:9092"
	asyncProducer, err := NewAsyncProducer(brokers)
	if err != nil {
		t.Fatal("TestNewAsyncProducer_case2 err == nil")
	}
	if asyncProducer == nil {
		t.Fatal("TestNewAsyncProducer_case2 asyncProducer != nil")
	}
	defer asyncProducer.Close()

	topic := "k8s-log-test-output-stdout"
	key := "TestSyncProducer_Send"
	value := "{\"name\":\"cgl\", \"message\":\"TestAsyncProducer_Send\"}"

	asyncProducer.Send(topic, key, value)
	time.Sleep(time.Second * 5)
}
