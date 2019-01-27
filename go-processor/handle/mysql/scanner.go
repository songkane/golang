// Package mysql scanner
// Created by chenguolin 2019-01-13
package mysql

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"gitlab.local.com/golang/go-log"
	"gitlab.local.com/golang/go-mysql"
	"gitlab.local.com/golang/go-processor/processor"
)

// Scanner mysql scanner
// chan T     send and receive type T data
// chan<- T   only send type T data
// <-chan int only receive type T data
type Scanner struct {
	isRunning    bool                  //scanner is running
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
		isRunning:    false,
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
	if s.isRunning {
		return
	}

	// set scanner running
	s.isRunning = true

	// new ticker
	// don't not use time.Timer, because Timer run only once
	ticker := time.Tick(s.scanInterval)

	// start scanner db
	go func() {
		defer func() {
			if err := recover(); err != nil {
				golog.Error("Mysql scanner start panic",
					golog.Object("Error", err))
				debug.PrintStack()
				// stop scanner
				s.Stop()
			}
		}()

		for {
			// stop return straight
			if !s.isRunning {
				return
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
	}()
}

// Stop scanner
func (s *Scanner) Stop() {
	if !s.isRunning {
		return
	}

	s.isRunning = false
	s.stopChan <- true
}

// Next get next record
func (s *Scanner) Next() (processor.Record, bool) {
	if !s.isRunning {
		return nil, false
	}

	// get record
	record, ok := <-s.records
	return record, ok
}

// IsStopped check scanner is running
func (s *Scanner) IsStopped() bool {
	return !s.isRunning
}
