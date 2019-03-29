// Package kafka scanner
// Created by chenguolin 2010-01-13
package kafka

import (
	"testing"
	"time"
)

func TestNewKafkaScanner(t *testing.T) {
	cfg := &Config{
		Topic:   "kafka_topic_test",
		GroupID: "consumer_group_id",
		Zk:      []string{"127.0.0.1:2181"},
	}
	scanner := NewKafkaScanner(cfg, 5)
	if scanner == nil {
		t.Fatal("TestNewKafkaScanner scanner == nil")
	}
	if scanner.IsRunning() == true {
		t.Fatal("TestNewKafkaScanner scanner.IsRunning() == true")
	}
	if scanner.maxChanSize != 5 {
		t.Fatal("TestNewKafkaScanner scanner.maxChanSize != 5")
	}
	if scanner.conf == nil {
		t.Fatal("TestNewKafkaScanner scanner.conf == nil")
	}
	if scanner.consumer == nil {
		t.Fatal("TestNewKafkaScanner scanner.consumer == nil")
	}
	if len(scanner.stopChan) != 0 {
		t.Fatal("TestNewKafkaScanner len(scanner.stopChan) != 0")
	}
	if len(scanner.records) != 0 {
		t.Fatal("TestNewKafkaScanner len(scanner.records) != 0")
	}
}

func TestScanner_Start(t *testing.T) {
	cfg := &Config{
		Topic:   "kafka_topic_test",
		GroupID: "consumer_group_id",
		Zk:      []string{"127.0.0.1:2181"},
	}
	scanner := NewKafkaScanner(cfg, 5)
	scanner.Start()
	if scanner.IsRunning() != true {
		t.Fatal("TestNewKafkaScanner scanner.IsRunning() != true")
	}
	if scanner.consumer.IsStopped() == true {
		t.Fatal("TestNewKafkaScanner scanner.consumer.IsStopped() == true")
	}
}

func TestScanner_Stop(t *testing.T) {
	cfg := &Config{
		Topic:   "kafka_topic_test",
		GroupID: "consumer_group_id",
		Zk:      []string{"127.0.0.1:2181"},
	}
	scanner := NewKafkaScanner(cfg, 5)
	scanner.Start()
	if scanner.IsRunning() == false {
		t.Fatal("TestNewKafkaScanner scanner.IsRunning == false")
	}
	// must be sleep wait consumer goroutine start
	time.Sleep(1 * time.Second)
	scanner.Stop()
	if scanner.IsStopped() == false {
		t.Fatal("TestNewKafkaScanner scanner.IsStopped() == false")
	}
}

func TestScanner_Next(t *testing.T) {
	cfg := &Config{
		Topic:   "kafka_topic_test",
		GroupID: "consumer_group_id",
		Zk:      []string{"127.0.0.1:2181"},
	}
	scanner := NewKafkaScanner(cfg, 5)
	scanner.Start()
	if scanner.IsStopped() == true {
		t.Fatal("TestScanner_Next scanner.IsStopped() == true")
	}
	time.Sleep(1 * time.Second)

	// not found record
	/*
		record, ok := scanner.Next()
		if ok == true {
			t.Fatal("TestScanner_Next ok == true")
		}
		fmt.Println(record)
	*/
}

func TestScanner_IsStopped(t *testing.T) {
	cfg := &Config{
		Topic:   "kafka_topic_test",
		GroupID: "consumer_group_id",
		Zk:      []string{"127.0.0.1:2181"},
	}
	scanner := NewKafkaScanner(cfg, 5)
	scanner.Start()
	if scanner.IsStopped() == true {
		t.Fatal("TestScanner_IsStopped scanner.IsStopped() == true")
	}
	time.Sleep(1 * time.Second)
	scanner.Stop()
	if scanner.IsStopped() == false {
		t.Fatal("TestScanner_IsStopped scanner.IsStopped() == false")
	}
}
