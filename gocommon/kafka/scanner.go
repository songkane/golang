// Package kafka kafka 消息scanner
// Created by chenguolin 2018-11-17
package kafka

import (
	"runtime/debug"

	"github.com/Shopify/sarama"
	cg "github.com/meitu/go-consumergroup"
	"gitlab.local.com/golang/golog"
	"gitlab.local.com/golang/httpserver/config"
)

const (
	ready    = 0
	running  = 1
	stopping = 3
	stopped  = 4
)

// Scanner kafka消息扫描器
type Scanner struct {
	topic    string
	consumer *cg.ConsumerGroup
	state    int
	buffer   <-chan *sarama.ConsumerMessage
}

// NewKafkaScanner 实例化一个kafka扫描器
func NewKafkaScanner(kafkaConf *config.KafkaConf, topic, consumerGroupID string) (*Scanner, error) {
	cgConf := cg.NewConfig()
	cgConf.ZkList = kafkaConf.ZK
	cgConf.TopicList = []string{topic}
	cgConf.GroupID = consumerGroupID
	consumer, err := cg.NewConsumerGroup(cgConf)
	if err != nil {
		return nil, err
	}
	return &Scanner{
		topic:    topic,
		consumer: consumer,
		state:    ready,
	}, nil
}

// Start 启动扫描器
func (scanner *Scanner) Start() {
	if scanner.state == running || scanner.state == stopping {
		return
	}
	err := scanner.consumer.Start()
	if err != nil {
		panic(err)
	}
	buf, _ := scanner.consumer.GetMessages(scanner.topic)
	scanner.buffer = buf
	scanner.state = running
	go func() {
		defer func() {
			if err := recover(); err != nil {
				golog.Error("kafka consumer group receive errors panic",
					golog.Object("error", err))
				debug.PrintStack()
			}
		}()
		errs, _ := scanner.consumer.GetErrors(scanner.topic)
		for {
			err, ok := <-errs
			if ok {
				if err != nil {
					golog.Error("kafka consumer group error",
						golog.Object("error", err))
				} else {
					golog.Error("kafka consumer group error,but received err is nil !!!")
				}
			} else {
				break
			}
		}
	}()
}

// Stop 停止扫描器
func (scanner *Scanner) Stop() {
	scanner.state = stopping
	scanner.consumer.Stop()
	scanner.state = stopped
}

// Next 获取下一个消息
func (scanner *Scanner) Next() (interface{}, bool) {
	v, ok := <-scanner.buffer
	return v, ok
}

func (scanner *Scanner) isStopped() bool {
	return scanner.state == stopped
}

func (scanner *Scanner) isStopping() bool {
	return scanner.state == stopping
}

func (scanner *Scanner) isRunning() bool {
	return scanner.state == running
}
