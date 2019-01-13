// Package mysql scanner
// Created by chenguolin 2019-01-13
package mysql

import (
	"time"

	"runtime/debug"

	"fmt"

	"gitlab.local.com/golang/golog"
	"gitlab.local.com/golang/gomysql"
	"gitlab.local.com/golang/goprocessor/processor"
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
	records      chan processor.Record //record channel
	dbProxy      *mysql.Mysql          //mysql proxy
}

// NewScanner new mysql scanner
func NewScanner(maxChanSize int, scanInterval time.Duration, dbProxy *mysql.Mysql) *Scanner {
	return &Scanner{
		isRunning:    false,
		maxChanSize:  maxChanSize,
		scanInterval: scanInterval,
		stopChan:     make(chan bool),
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
			}
		}()

		for {
			if !s.isRunning {
				return
			}

			select {
			case <-s.stopChan:
				// close records channel
				close(s.records)
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
