// Package kafka scanner
// Created by chenguolin 2010-01-13
package kafka

import (
	"runtime/debug"

	"github.com/Shopify/sarama"
	cg "github.com/meitu/go-consumergroup"
	"gitlab.local.com/golang/golog"
	"gitlab.local.com/golang/goprocessor/processor"
)

// Scanner kafka scanner
// chan T     send and receive type T data
// chan<- T   only send type T data
// <-chan int only receive type T data
type Scanner struct {
	isRunning   bool                           //scanner is running
	maxChanSize int                            //max channel size
	conf        *Config                        //kafka config
	consumer    *cg.ConsumerGroup              //kafka consumer
	stopChan    chan bool                      //stop channel
	records     <-chan *sarama.ConsumerMessage //record channel
}

// Config kafka
// Read kafka only need Zk
// Write kafka only need brokers
type Config struct {
	Topic   string   `json:"topic"`    //topic
	GroupID string   `json:"group_id"` //consumer group id
	Zk      []string `json:"zk"`       //zk
}

// NewKafkaScanner new kafka scanner
func NewKafkaScanner(conf *Config, maxChanSize int) *Scanner {
	if conf == nil || maxChanSize <= 0 {
		panic("NewKafkaScanner panic")
	}

	// new consumer group config
	cgConf := cg.NewConfig()
	cgConf.ZkList = conf.Zk
	cgConf.TopicList = []string{conf.Topic}
	cgConf.GroupID = conf.GroupID

	// new consumer group
	consumer, err := cg.NewConsumerGroup(cgConf)
	if err != nil {
		panic(err)
	}

	return &Scanner{
		isRunning:   false,
		maxChanSize: maxChanSize,
		conf:        conf,
		consumer:    consumer,
		stopChan:    make(chan bool),
		records:     make(chan *sarama.ConsumerMessage, maxChanSize),
	}
}

// Start scanner
func (s *Scanner) Start() {
	if s.isRunning {
		return
	}

	// set running true
	s.isRunning = true

	// start consumer
	err := s.consumer.Start()
	if err != nil {
		panic(err)
	}

	// go start
	go func() {
		defer func() {
			if err := recover(); err != nil {
				golog.Error("Kafka scanner start panic",
					golog.Object("Error", err))
				debug.PrintStack()
				// stop scanner
				s.Stop()
			}
		}()

		// get message channel
		msgChan, ok := s.consumer.GetMessages(s.conf.Topic)
		if !ok {
			golog.Panic("consumer.GetMessage not successful")
		}
		s.records = msgChan
	}()
}

// Stop scanner
func (s *Scanner) Stop() {
	if !s.isRunning {
		return
	}

	// stop consumer
	s.isRunning = false
	s.consumer.Stop()
}

// Next get next kafka message
func (s *Scanner) Next() (processor.Record, bool) {
	if !s.isRunning {
		return nil, false
	}

	// next record
	record, ok := <-s.records
	return record, ok
}

// IsStopped check scanner is running
func (s *Scanner) IsStopped() bool {
	return !s.isRunning
}
