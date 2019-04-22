// Package kafka sync producer unit test
// Created by chenguolin 2019-04-20
package kafka

import (
	"testing"
)

func TestNewSyncProducer_case1(t *testing.T) {
	syncProducer, err := NewSyncProducer("")
	if err == nil {
		t.Fatal("TestNewSyncProducer_case1 err == nil")
	}
	if syncProducer != nil {
		t.Fatal("TestNewSyncProducer_case1 syncProducer != nil")
	}
}

func TestNewSyncProducer_case2(t *testing.T) {
	brokers := "172.16.28.47:9092,172.16.28.48:9092,172.16.28.49:9092"
	syncProducer, err := NewSyncProducer(brokers)
	if err != nil {
		t.Fatal("TestNewSyncProducer_case2 err != nil")
	}
	if syncProducer == nil {
		t.Fatal("TestNewSyncProducer_case2 syncProducer == nil")
	}
	defer syncProducer.Close()
}

func TestSyncProducer_Send(t *testing.T) {
	brokers := "172.16.28.47:9092,172.16.28.48:9092,172.16.28.49:9092"
	syncProducer, err := NewSyncProducer(brokers)
	if err != nil {
		t.Fatal("TestSyncProducer_Send err != nil")
	}
	if syncProducer == nil {
		t.Fatal("TestSyncProducer_Send syncProducer == nil")
	}
	defer syncProducer.Close()

	topic := "k8s-log-test-output-stdout"
	key := "TestSyncProducer_Send"
	value := "{\"name\":\"cgl\", \"message\":\"TestSyncProducer_Send\"}"

	partition, offset, err := syncProducer.Send(topic, key, value)
	if partition != 0 {
		t.Fatal("TestSyncProducer_Send partition != 0")
	}
	if offset < 0 {
		t.Fatal("TestSyncProducer_Send offset < 0")
	}
	if err != nil {
		t.Fatal("TestSyncProducer_Send err != nil")
	}
}
