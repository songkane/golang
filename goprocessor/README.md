# goprocessor
golang processor

1. processor 包含3个go文件
   * processor: Processor struct
   * scanner: scanner interface
   * handle: handle interface
2. handle/mysql mysql processor
   * scanner: mysql scanner
   * handle: mysql handle
3. handle/kafka kafka processor
   * scanner: kafka scanner
   * handle: kafka handle

如果要实现新的processor只需要在handle目录下参考现有的mysql、kafka实现scanner和handle即可

# sample
```
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
    Topic:           "test_topic",
	ConsumerGroupId: "test_consumer_group",
	Zk:              []string{"127.0.0.1:2181"},
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
```
