// Package kafka consumer
// Created by chenguolin 2019-04-20
package kafka

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.local.com/golang/go-kafka/pkg/sarama"
	cluster "gitlab.local.com/golang/go-kafka/pkg/sarama-cluster"
)

// Error define
var (
	ErrBrokersIsEmpty = errors.New("brokers is empty")
	ErrTopicIsEmpty   = errors.New("topic is empty")
	ErrGroupIDIsEmpty = errors.New("groupId is empty")
)

// Offset define
const (
	OffsetNewset = iota
	OffsetOldest
)

// Consumer client
type Consumer struct {
	topic         string            //topic
	groupID       string            //group id
	brokers       string            //brokers
	defaultOffset int               //default offset
	consumer      *cluster.Consumer //consumer group
}

// NewConsumer new consumer
func NewConsumer(brokers, topic, groupID string, defaultOffset int) (*Consumer, error) {
	// check args
	if brokers == "" {
		return nil, ErrBrokersIsEmpty
	}
	if topic == "" {
		return nil, ErrTopicIsEmpty
	}
	if groupID == "" {
		return nil, ErrGroupIDIsEmpty
	}

	// new config
	// default read from OffsetNewest
	cfg := cluster.NewConfig()
	// return errors
	cfg.Consumer.Return.Errors = true

	// 如果zk已经存在当前consumer group，则从上次中断的地方开始消费，例如程序重启
	// 如果zk不存在当前consumer group，通过defaultOffset设置从最老还是最新开始消费
	// 只允许设置为sarama.OffsetOldest 或 sarama.OffsetNewest
	if defaultOffset == OffsetOldest {
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	} else {
		cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	// TODO set tls config (@cgl)
	// if *tlsEnabled {
	// 		tlsConfig, err := tls.NewConfig(*tlsClientCert, *tlsClientKey)
	// 		if err != nil {
	// 			printErrorAndExit(69, "Failed to create TLS config: %s", err)
	// 		}
	//
	// 		config.Net.TLS.Enable = true
	// 		config.Net.TLS.Config = tlsConfig
	// 		config.Net.TLS.Config.InsecureSkipVerify = *tlsSkipVerify
	// }

	// new consumer group
	topics := []string{topic}
	consumer, err := cluster.NewConsumer(strings.Split(brokers, ","), groupID, topics, cfg)
	if err != nil {
		return nil, fmt.Errorf("cluster.NewConsumer error, %s", err.Error())
	}

	// return consumer
	c := &Consumer{
		topic:         topic,
		groupID:       groupID,
		brokers:       brokers,
		defaultOffset: defaultOffset,
		consumer:      consumer,
	}

	return c, nil
}

// Messages return message channel
func (c *Consumer) Messages() <-chan *sarama.ConsumerMessage {
	return c.consumer.Messages()
}

// Errors return errors channel
func (c *Consumer) Errors() <-chan error {
	return c.consumer.Errors()
}

// CommitOffset mark message as processed
func (c *Consumer) CommitOffset(msg *sarama.ConsumerMessage) {
	c.consumer.MarkOffset(msg, "")
}

// Close consumer
func (c *Consumer) Close() {
	c.consumer.Close()
}
