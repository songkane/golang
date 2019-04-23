// Package kafka get metadata
// Created by chenguolin 2019-04-20
package kafka

import (
	"errors"
	"strings"

	"github.com/Shopify/sarama"
)

// Client is a generic Kafka client. It manages connections to one or more Kafka brokers.
// You MUST call Close() on a client to avoid leaks, it will not be garbage-collected
// automatically when it passes out of scope. It is safe to share a client amongst many
// users, however Kafka will process requests from a single client strictly in serial,
// so it is generally more efficient to use the default one client per producer/consumer.
type Client struct {
	brokers string        //brokers
	client  sarama.Client //client
}

// NewClient new kafka client
func NewClient(brokers string) (*Client, error) {
	if brokers == "" {
		return nil, errors.New("brokers is empty")
	}

	cfg := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(brokers, ","), cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		brokers: brokers,
		client:  client,
	}, nil
}

// Brokers returns the current set of active brokers as retrieved from cluster metadata.
func (c *Client) Brokers() []*sarama.Broker {
	return c.client.Brokers()
}

// Topics returns the set of available topics as retrieved from cluster metadata
func (c *Client) Topics() ([]string, error) {
	return c.client.Topics()
}

// Partitions returns the sorted list of all partition IDs for the given topic.
func (c *Client) Partitions(topic string) ([]int32, error) {
	return c.client.Partitions(topic)
}

// WritablePartitions returns the sorted list of all writable partition IDs for
// the given topic, where "writable" means "having a valid leader accepting
// writes".
func (c *Client) WritablePartitions(topic string) ([]int32, error) {
	return c.client.WritablePartitions(topic)
}

// Leader returns the broker object that is the leader of the current
// topic/partition, as determined by querying the cluster metadata.
func (c *Client) Leader(topic string, partitionID int32) (*sarama.Broker, error) {
	return c.client.Leader(topic, partitionID)
}

// Replicas returns the set of all replica IDs for the given partition.
func (c *Client) Replicas(topic string, partitionID int32) ([]int32, error) {
	return c.client.Replicas(topic, partitionID)
}

// GetOffset queries the cluster to get the most recent available offset at the
// given time (in milliseconds) on the topic/partition combination.
// Time should be OffsetOldest or OffsetNewest or a time.
func (c *Client) GetOffset(topic string, partitionID int32, time int64) (int64, error) {
	return c.client.GetOffset(topic, partitionID, time)
}

// return a locally cached value if it's available. You can call
// RefreshCoordinator to update the cached value. This function only works on
// Kafka 0.8.2 and higher.
func (c *Client) Coordinator(consumerGroup string) (*sarama.Broker, error) {
	return c.client.Coordinator(consumerGroup)
}
