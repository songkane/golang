// Package kafka sync producer
// Created by chenguolin 2019-04-20
package kafka

import (
	"strings"

	"github.com/Shopify/sarama"
)

// SyncProducer sync producer client
// SyncProducer publishes Kafka messages, blocking until they have been acknowledged. It routes messages to the correct
// broker, refreshing metadata as appropriate, and parses responses for errors. You must call Close() on a producer
// to avoid leaks, it may not be garbage-collected automatically when it passes out of scope.
//
// The SyncProducer comes with two caveats: it will generally be less efficient than the AsyncProducer, and the actual
// durability guarantee provided when a message is acknowledged depend on the configured value of `Producer.RequiredAcks`.
// There are configurations where a message acknowledged by the SyncProducer can still sometimes be lost.
//
// For implementation reasons, the SyncProducer requires `Producer.Return.Errors` and `Producer.Return.Successes` to
// be set to true in its configuration.
type SyncProducer struct {
	brokers  string              //brokers
	producer sarama.SyncProducer //producer
}

// NewSyncProducer new Producer client
func NewSyncProducer(brokers string) (*SyncProducer, error) {
	// check args
	if brokers == "" {
		return nil, ErrBrokersIsEmpty
	}

	// sarama config
	cfg := sarama.NewConfig()
	// WaitForAll waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.insync.replicas` configuration key.
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true

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

	// new sync producer
	producer, err := sarama.NewSyncProducer(strings.Split(brokers, ","), cfg)
	if err != nil {
		return nil, err
	}

	return &SyncProducer{
		brokers:  brokers,
		producer: producer,
	}, nil
}

// Send message 2 kafka
func (sp *SyncProducer) Send(topic, key, value string) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	return sp.producer.SendMessage(msg)
}

// Close producer
func (sp *SyncProducer) Close() error {
	return sp.producer.Close()
}
