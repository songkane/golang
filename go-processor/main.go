// Package main golang processor main entry
// Created by chenguolin 2019-01-13
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.local.com/golang/go-processor/handle/kafka"
	"gitlab.local.com/golang/go-processor/handle/mysql"
	"gitlab.local.com/golang/go-processor/processor"
)

func main() {
	// mysql processor
	maxChanSize := 50
	scanInterval := 5 * time.Second
	// dbProxy := db.NewMysql(nil)
	mysqlScanner := mysql.NewScanner(maxChanSize, scanInterval, nil)
	mysqlHandle := mysql.NewHandle()
	concurrentCnt := 4
	mysqlProcessor := processor.NewProcessor(mysqlScanner, mysqlHandle, concurrentCnt)
	mysqlProcessor.Start()
	fmt.Println("Mysql processor start ...")

	// kafka processor
	maxChanSize = 100
	kafkaConf := &kafka.Config{
		Topic:   "test_topic",
		GroupID: "test_consumer_group",
		Zk:      []string{"127.0.0.1:2181"},
	}
	kafkaScanner := kafka.NewKafkaScanner(kafkaConf, maxChanSize)
	kafkaHandle := kafka.NewHandle()
	concurrentCnt = 2
	kafkaProcessor := processor.NewProcessor(kafkaScanner, kafkaHandle, concurrentCnt)
	kafkaProcessor.Start()
	fmt.Println("Kafka processor start ...")

	// wait shutdown
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stopSignalChan
	fmt.Println("Got system signal:" + sig.String() + ", going to shutdown.")
	// stop mysql processor
	mysqlProcessor.Stop()
	fmt.Println("Mysql processor stop successful ~")
	// stop kafka processor
	kafkaProcessor.Stop()
	fmt.Println("Kafka processor stop successful ~")
}