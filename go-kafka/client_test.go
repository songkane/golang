// Package kafka get metadata
// Created by chenguolin 2019-04-20
package kafka

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestNewClient(t *testing.T) {
	// case 1
	brokers := ""
	client, err := NewClient(brokers)
	if err == nil {
		t.Fatal("TestNewClient case 1 err == nil")
	}
	if client != nil {
		t.Fatal("TestNewClient case 1 client != nil")
	}

	// case 2
	brokers = "192.168.0.1:9092,192.168.0.2:9092"
	client, err = NewClient(brokers)
	if err != nil {
		t.Fatal("TestNewClient case 2 err != nil")
	}
	if client == nil {
		t.Fatal("TestNewClient case 2 client == nil")
	}
}

func TestClient_Brokers(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Brokers err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Brokers client == nil")
	}

	bks := client.Brokers()
	if bks == nil {
		t.Fatal("TestClient_Brokers bks != nil")
	}
	fmt.Println(bks)
}

func TestClient_Topics(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Topics err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Topics client == nil")
	}

	topics, err := client.Topics()
	if err != nil {
		t.Fatal("TestClient_Topics err != nil")
	}
	for _, topic := range topics {
		fmt.Println(topic)
	}
}

func TestClient_Partitions(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Partitions err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Partitions client == nil")
	}

	topic := "k8s-log-test-output-stdout"
	partitions, err := client.Partitions(topic)
	if err != nil {
		t.Fatal("TestClient_Partitions client.Partitions err != nil")
	}
	if len(partitions) <= 0 {
		t.Fatal("TestClient_Partitions len(partitions) <= 0")
	}
	fmt.Println(partitions)
}

func TestClient_WritablePartitions(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_WritablePartitions err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_WritablePartitions client == nil")
	}

	topic := "k8s-log-test-output-stdout"
	partitions, err := client.WritablePartitions(topic)

	if err != nil {
		t.Fatal("TestClient_WritablePartitions client.WritablePartitions err != nil")
	}
	if len(partitions) <= 0 {
		t.Fatal("TestClient_WritablePartitions len(partitions) <= 0")
	}
	fmt.Println(partitions)
}

func TestClient_Leader(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Leader err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Leader client == nil")
	}

	topic := "k8s-log-test-output-stdout"
	partitionID := int32(0)

	broker, err := client.Leader(topic, partitionID)
	if err != nil {
		t.Fatal("client.Leader err != nil")
	}
	fmt.Println(broker.Addr())
}

func TestClient_Replicas(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Leader err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Leader client == nil")
	}

	topic := "k8s-log-test-output-stdout"
	partitionID := int32(0)

	replicas, err := client.Replicas(topic, partitionID)
	if err != nil {
		t.Fatal("TestClient_Leader client.Replicas err != nil")
	}
	if len(replicas) <= 0 {
		t.Fatal("TestClient_Leader len(replicas) <= 0 ")
	}
	fmt.Println(replicas)
}

func TestClient_GetOffset(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_GetOffset err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_GetOffset client == nil")
	}

	topic := "k8s-log-test-output-stdout"
	partitionID := int32(0)

	// case 1
	offset, err := client.GetOffset(topic, partitionID, sarama.OffsetOldest)
	if err != nil {
		t.Fatal("TestClient_GetOffset case 1 client.GetOffset err != nil")
	}
	fmt.Println(offset)

	// case 2
	offset, err = client.GetOffset(topic, partitionID, sarama.OffsetNewest)
	if err != nil {
		t.Fatal("TestClient_GetOffset case 2 client.GetOffset err != nil")
	}
	fmt.Println(offset)

	// case 3
	offset, err = client.GetOffset(topic, partitionID, time.Now().Unix()*1000)
	if err != nil {
		t.Fatal("TestClient_GetOffset case 3 client.GetOffset err != nil")
	}
	fmt.Println(offset)
}

func TestClient_Coordinator(t *testing.T) {
	brokers := "192.168.0.1:9092,192.168.0.2:9092"
	client, err := NewClient(brokers)
	if err != nil {
		t.Fatal("TestClient_Coordinator err != nil")
	}
	if client == nil {
		t.Fatal("TestClient_Coordinator client == nil")
	}

	consumerGroup := "TestConsumer_Messages"
	broker, err := client.Coordinator(consumerGroup)
	if err != nil {
		t.Fatal("TestClient_Coordinator client.Coordinator err != nil")
	}

	// get offset
	request := &sarama.ConsumerMetadataRequest{
		ConsumerGroup: consumerGroup,
	}

	resp, err := broker.GetConsumerMetadata(request)
	fmt.Println(err)
	bytes, _ := json.Marshal(resp)
	fmt.Println(string(bytes))
}
