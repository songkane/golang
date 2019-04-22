// Package kafka 通用写kafka接口
// Created by chenguolin 2018-09-21
package kafka

import (
	"errors"
	"time"

	"gitlab.local.com/golang/go-kafka"
	golog "gitlab.local.com/golang/go-log"
)

// NewSyncKafkaProducer 实例化一个kafka producer
func NewSyncKafkaProducer(brokers string) (*kafka.SyncProducer, error) {
	producer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

// WriteKafka 写Kafka消息接口
// kafkaProducer 生产者对象
// topic 要写的topic
// key 消息key
// value 消息内容
// retryTimes 重试次数
func WriteKafka(producer *kafka.SyncProducer, topic, key, value string, retryTimes int) error {
	tryTimes := 0

	for {
		if tryTimes >= retryTimes {
			break
		}
		tryTimes++

		// 写消息
		partition, offset, err := producer.Send(topic, key, value)
		if err != nil {
			golog.Warn("WriteKafka write 2 kakfa failed.",
				golog.Int("retryTimes", tryTimes),
				golog.String("topic", topic),
				golog.String("key", key),
				golog.String("value", value),
				golog.Object("error", err))
			time.Sleep(time.Duration(2*tryTimes) * time.Second)
		} else {
			golog.Info("WriteKafka write 2 kafka sucessful ~",
				golog.String("topic", topic),
				golog.Object("partition", partition),
				golog.Object("offset", offset))
			return nil
		}
	}

	return errors.New("WriteKafka write 2 kafka failed ~ " + "topic: " + topic)
}
