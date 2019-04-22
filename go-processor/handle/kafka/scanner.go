// Package kafka scanner
// Created by chenguolin 2010-01-13
package kafka

import (
	"github.com/Shopify/sarama"
	"gitlab.local.com/golang/go-kafka"
	"gitlab.local.com/golang/go-processor/processor"
)

// scanner state
const (
	running = iota
	stopped
)

// Scanner kafka scanner
// chan T     send and receive type T data
// chan<- T   only send type T data
// <-chan T   only receive type T data
type Scanner struct {
	state       int                            //scanner state
	maxChanSize int                            //max channel size
	consumer    *kafka.Consumer                //kafka consumer
	stopChan    chan bool                      //stop channel
	msgChan     <-chan *sarama.ConsumerMessage //message channel
}

// NewKafkaScanner new kafka scanner
func NewKafkaScanner(brokers, topic, groupID string, maxChanSize int) *Scanner {
	if brokers == "" || topic == "" || groupID == "" || maxChanSize <= 0 {
		panic("NewKafkaScanner panic, invalid args")
	}

	// new consumer group
	consumer, err := kafka.NewConsumer(brokers, topic, groupID, kafka.OffsetNewset)
	if err != nil {
		panic(err)
	}

	return &Scanner{
		state:       running,
		maxChanSize: maxChanSize,
		consumer:    consumer,
		stopChan:    make(chan bool),
		msgChan:     make(<-chan *sarama.ConsumerMessage, maxChanSize),
	}
}

// Start scanner
func (s *Scanner) Start() {
	if s.state == running {
		return
	}

	// set state 2 running
	s.state = running
	s.msgChan = s.consumer.Messages()
}

// Stop scanner
func (s *Scanner) Stop() {
	if s.state != running {
		return
	}

	// set state 2 stopped
	s.state = stopped
	// close consumer
	s.consumer.Close()
}

// Next get next kafka message
func (s *Scanner) Next() (processor.Record, bool) {
	// 1. channel关闭后，继续往它发送数据会panic
	// 2. channel关闭后，使用 <-c 方式还可以继续读到数据，只不过读到的是对应类型的零值
	// 3. 通过for range的方式读取channel，channel关闭后会退出for循环
	// 4. 通过 v,ok := <- c方式读取channel, 如果channel没有数据或者channel关闭，v为对应类型的零值，ok为false
	//    (注意不能简单的通过ok为false来判断channel已经关闭，因为有可能是channel没有数据)

	// next record
	record, ok := <-s.msgChan
	return record, ok
}

// IsRunning check scanner is running
func (s *Scanner) IsRunning() bool {
	return s.state == running
}

// IsStopped check scanner is stopped
func (s *Scanner) IsStopped() bool {
	return s.state == stopped
}
