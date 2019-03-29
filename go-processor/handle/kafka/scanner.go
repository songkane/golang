// Package kafka scanner
// Created by chenguolin 2010-01-13
package kafka

import (
	"runtime/debug"

	"github.com/Shopify/sarama"
	cg "github.com/meitu/go-consumergroup"

	golog "gitlab.local.com/golang/go-log"
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
		state:       running,
		maxChanSize: maxChanSize,
		conf:        conf,
		consumer:    consumer,
		stopChan:    make(chan bool),
		records:     make(chan *sarama.ConsumerMessage, maxChanSize),
	}
}

// Start scanner
func (s *Scanner) Start() {
	if s.state == running {
		return
	}

	// set state 2 running
	s.state = running

	// start consumer
	err := s.consumer.Start()
	if err != nil {
		panic(err)
	}

	// start scan
	go s.scan()
}

// scan records
func (s *Scanner) scan() {
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
}

// Stop scanner
func (s *Scanner) Stop() {
	if s.state != running {
		return
	}

	// set state 2 stopped
	s.state = stopped
	// stop consumer
	s.consumer.Stop()
}

// Next get next kafka message
func (s *Scanner) Next() (processor.Record, bool) {
	// 1. channel关闭后，继续往它发送数据会panic
	// 2. channel关闭后，使用 <-c 方式还可以继续读到数据，只不过读到的是对应类型的零值
	// 3. 通过for range的方式读取channel，channel关闭后会退出for循环
	// 4. 通过 v,ok := <- c方式读取channel, 如果channel没有数据或者channel关闭，v为对应类型的零值，ok为false
	//    (注意不能简单的通过ok为false来判断channel已经关闭，因为有可能是channel没有数据)

	// next record
	record, ok := <-s.records
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
