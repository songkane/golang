// Package kafka async producer
// Created by chenguolin 2019-04-20
package kafka

import (
	"runtime/debug"
	"strings"
	"sync"

	"gitlab.local.com/golang/go-kafka/pkg/sarama"
	golog "gitlab.local.com/golang/go-log"
)

// AsyncProducer async producer client
// AsyncProducer publishes Kafka messages using a non-blocking API. It routes messages
// to the correct broker for the provided topic-partition, refreshing metadata as appropriate,
// and parses responses for errors. You must read from the Errors() channel or the
// producer will deadlock. You must call Close() or AsyncClose() on a producer to avoid
// leaks: it will not be garbage-collected automatically when it passes out of
// scope.
type AsyncProducer struct {
	brokers  string               //brokers
	wg       *sync.WaitGroup      //waitGroup
	stopChan chan struct{}        //stop channel
	producer sarama.AsyncProducer //producer
}

// NewAsyncProducer new Producer client
func NewAsyncProducer(brokers string) (*AsyncProducer, error) {
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
	producer, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), cfg)
	if err != nil {
		return nil, err
	}

	ap := &AsyncProducer{
		brokers:  brokers,
		wg:       &sync.WaitGroup{},
		stopChan: make(chan struct{}),
		producer: producer,
	}

	// goroutine receive error chanel
	ap.wg.Add(1)
	go ap.receiveError()

	return ap, nil
}

// Send message 2 kafka
func (ap *AsyncProducer) Send(topic, key, value string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	if key != "" {
		msg.Key = sarama.StringEncoder(key)
	}

	// send 2 input channel
	ap.producer.Input() <- msg
}

// Close async producer
func (ap *AsyncProducer) Close() error {
	err := ap.producer.Close()
	if err != nil {
		return err
	}

	// close stopChan
	close(ap.stopChan)
	ap.wg.Wait()

	return nil
}

// receiveError receive from Success and Errors channel
func (ap *AsyncProducer) receiveError() {
	defer func() {
		if err := recover(); err != nil {
			golog.Error("AsyncProducer receiveError handle panic",
				golog.Object("error", err))
			debug.PrintStack()
		}
		// goroutine done
		ap.wg.Done()
	}()

	for {
		select {
		case err := <-ap.producer.Errors():
			if err != nil {
				golog.Error("AsyncProducer receive error",
					golog.Object("error", err))
			}
		case <-ap.stopChan:
			return
		}
	}
}
