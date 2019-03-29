// Package mysql scanner
// Created by chenguolin 2019-01-13
package mysql

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	golog "gitlab.local.com/golang/go-log"
	"gitlab.local.com/golang/go-mysql"
	"gitlab.local.com/golang/go-processor/processor"
)

// scanner state
const (
	running = iota
	stopped
)

// Scanner mysql scanner
// chan T     send and receive type T data
// chan<- T   only send type T data
// <-chan T   only receive type T data
type Scanner struct {
	state        int                   //scanner state
	maxChanSize  int                   //max records channel size
	scanInterval time.Duration         //scanner interval
	stopChan     chan bool             //stop channel
	closeOnce    sync.Once             //stop only once
	records      chan processor.Record //record channel
	dbProxy      *mysql.Mysql          //mysql proxy
}

// NewScanner new mysql scanner
func NewScanner(maxChanSize int, scanInterval time.Duration, dbProxy *mysql.Mysql) *Scanner {
	if maxChanSize <= 0 || dbProxy == nil {
		panic("NewScanner panic")
	}

	return &Scanner{
		state:        stopped,
		maxChanSize:  maxChanSize,
		scanInterval: scanInterval,
		stopChan:     make(chan bool),
		closeOnce:    sync.Once{},
		records:      make(chan processor.Record, maxChanSize),
		dbProxy:      dbProxy,
	}
}

// Start scanner
func (s *Scanner) Start() {
	if s.state == running {
		return
	}

	// set scanner running
	s.state = running

	// start scan
	go s.scan()
}

// scan records
func (s *Scanner) scan() {
	// start scanner db
	defer func() {
		if err := recover(); err != nil {
			golog.Error("Mysql scanner start panic",
				golog.Object("Error", err))
			debug.PrintStack()
			// stop scanner
			s.Stop()
		}
	}()

	// new ticker
	// don't not use time.Timer, because Timer run only once
	ticker := time.Tick(s.scanInterval)

	for {
		// 如果已经是stopped状态直接退出
		if s.state == stopped {
			break
		}

		select {
		case <-s.stopChan:
			// sync.Once能确保实例化对象Do方法在多线程环境只运行一次
			// 防止多次close导致panic

			// close records channel
			s.closeOnce.Do(func() {
				close(s.records)
			})
			return
		case <-ticker:
			// scan mysql table
			// TODO select from mysql db
			// TODO insert record 2 records channel
			// TODO rds := dbProxy.Query("select ...")
			// TODO for _, record := range rds {
			// TODO     s.records <- record
			// TODO }

			record := fmt.Sprintf("%s: mysql record", time.Now().Format("2006-01-02 15:04:05"))
			s.records <- record
		}
	}
}

// Stop scanner
func (s *Scanner) Stop() {
	if s.state != running {
		return
	}

	// set stopped
	s.state = stopped
	// write 2 stopChan
	s.stopChan <- true
}

// Next get next record
func (s *Scanner) Next() (processor.Record, bool) {
	// 1. channel关闭后，继续往它发送数据会panic
	// 2. channel关闭后，使用 <-c 方式还可以继续读到数据，只不过读到的是对应类型的零值
	// 3. 通过for range的方式读取channel，channel关闭后会退出for循环
	// 4. 通过 v,ok := <- c方式读取channel, 如果channel没有数据或者channel关闭，v为对应类型的零值，ok为false
	//    (注意不能简单的通过ok为false来判断channel已经关闭，因为有可能是channel没有数据)

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
