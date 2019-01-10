// Package kafka 通用写kafka接口
// Created by chenguolin 2018-09-21
package kafka

import (
	"errors"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.local.com/golang/golog"
)

// NewSyncKafkaProducer 实例化一个kafka producer
func NewSyncKafkaProducer(brokers []string) (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer, nil
}

// WriteKafka 写Kafka消息接口
// kafkaProducer 生产者对象
// msg 消息
// retryTimes 重试次数
func WriteKafka(kafkaProducer sarama.SyncProducer, msg *sarama.ProducerMessage, retryTimes int) error {
	tryTimes := 0

	for {
		if tryTimes >= retryTimes {
			break
		}
		tryTimes++

		// 写消息
		partition, offset, err := kafkaProducer.SendMessage(msg)
		if err != nil {
			golog.Warn("[kafka - WriteKafka] write 2 kakfa failed.",
				golog.Int("retryTimes", tryTimes),
				golog.Object("message", msg),
				golog.Object("Error", err))
			time.Sleep(time.Duration(2*tryTimes) * time.Second)
		} else {
			golog.Info("[kafka - WriteKafka] write 2 kafka sucessful ~",
				golog.String("topic", msg.Topic),
				golog.Object("partition", partition),
				golog.Object("offset", offset),
				golog.Object("message", msg))
			return nil
		}
	}

	return errors.New("[kafka - WriteKafka] write 2 kafka failed ~ " + "topic: " + msg.Topic)
}
