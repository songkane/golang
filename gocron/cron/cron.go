// Package cron golang crontab
// Created by chenguolin 2018-01-04
package cron

import (
	"runtime/debug"
	"sort"
	"time"

	"gitlab.local.com/golang/golog"
)

// Cron define
type Cron struct {
	entryList entrys    //cron entry list
	isRunning bool      //cron is running status
	stopChan  chan bool //cron stop channel
}

// NewCron new cron
func NewCron() *Cron {
	return &Cron{
		entryList: make([]*entry, 0),
		isRunning: false,
		stopChan:  make(chan bool),
	}
}

// AddHandle add new handle
func (c *Cron) AddHandle(scheduler *Scheduler, handle handle) {
	// 1. check scheduler
	if scheduler == nil {
		golog.Error("AddHandle scheduler is nil")
	}

	// 2. new entry
	entry := &entry{
		scheduler: scheduler,
		handle:    handle,
		nextTime:  scheduler.Next(time.Now()),
	}

	// 3. entry append 2 list
	c.entryList = append(c.entryList, entry)
}

// Start cron
// if cron already running no operator
func (c *Cron) Start() {
	if c.isRunning {
		return
	}

	if len(c.entryList) <= 0 {
		return
	}

	// set isRunning true
	c.isRunning = true
	go c.run()
}

// run cron
func (c *Cron) run() {
	defer func() {
		if r := recover(); r != nil {
			golog.Error("cron run panic recover")
			debug.PrintStack()
		}
	}()

	for {
		if !c.isRunning {
			return
		}

		// 1. sort all entry by NextTime
		sort.Sort(entrys(c.entryList))

		// 2. new timer
		// use time.Timer, because Timer run only once
		timer := time.NewTimer(c.entryList[0].nextTime.Sub(time.Now()))

		// 3. handle handleProcess
		select {
		case now := <-timer.C:
			// foreach entry
			for _, e := range c.entryList {
				if e.nextTime.After(now) {
					break
				}
				// goroutine call handleProcess
				go c.handleProcess(e.handle)
				e.preTime = now
				e.nextTime = e.scheduler.Next(now)
			}
		case <-c.stopChan:
			// timer stop
			timer.Stop()
			// return function
			// don't not use break
			return
		}
	}
}

// handleProcess call handle DoProcess function
func (c *Cron) handleProcess(handle handle) {
	defer func() {
		if r := recover(); r != nil {
			golog.Error("cron handleProcess panic recover")
			debug.PrintStack()
		}
	}()

	// call handle DoProcess function
	handle.DoProcess()
}

// Stop cron
// if cron already stop no operator
func (c *Cron) Stop() {
	if !c.isRunning {
		return
	}

	// set is running 2 false
	c.isRunning = false
	// write 2 stop channel
	c.stopChan <- true
}
