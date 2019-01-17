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
	entries  []*entry  // entries: all run Entry set
	running  bool      // running: crontab is running
	stopChan chan bool // stopChan: crontab stop channel
}

// entry define
type entry struct {
	scheduler *Scheduler // Schedule: job
	job       FuncJob    // Job: run function
	nextTime  time.Time  // NextTime: next run time
	preTime   time.Time  // PreTime: last run time
}

// FuncJob func helper type
type FuncJob func()

// NewCron crontab
// default use time location
func NewCron() *Cron {
	return &Cron{
		entries:  make([]*entry, 0),
		running:  false,
		stopChan: make(chan bool),
	}
}

// AddJob add new job 2 crontab
func (c *Cron) AddJob(sche *Scheduler, job FuncJob) {
	// 1. check args
	if sche == nil || job == nil {
		golog.Error("AddJob scheduler or job invalid")
	}

	// 2. add 2 entries
	entry := &entry{
		scheduler: sche,
		job:       job,
		nextTime:  sche.Next(time.Now()),
	}
	c.entries = append(c.entries, entry)
}

// entrys []*entry helper struct
type entrys []*entry

func (es entrys) Len() int {
	return len(es)
}

func (es entrys) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es entrys) Less(i, j int) bool {
	if es[i].nextTime.IsZero() {
		return false
	}
	if es[j].nextTime.IsZero() {
		return true
	}

	return es[i].nextTime.Before(es[j].nextTime)
}

// Start crontab
// if crontab already running no operator
func (c *Cron) Start() {
	if c.running {
		return
	}
	if len(c.entries) <= 0 {
		return
	}

	// set running true
	c.running = true
	go c.run()
}

func (c *Cron) run() {
	defer func() {
		if r := recover(); r != nil {
			golog.Error("cron run panic recover")
			debug.PrintStack()
		}
	}()

	for {
		if !c.running {
			return
		}

		// 1. sort all entry by NextTime
		sort.Sort(entrys(c.entries))

		// 2. new timer
		// everytime need new Timer, because Timer run only once
		timer := time.NewTimer(c.entries[0].nextTime.Sub(time.Now()))

		// 3. start run job
		// process label
		for {
			select {
			case now := <-timer.C:
				// run all has same NextTime entry
				for _, e := range c.entries {
					if e.nextTime.After(now) {
						break
					}
					// goroutine
					go c.process(e.job)
					e.preTime = now
					e.nextTime = e.scheduler.Next(now)
				}
			case <-c.stopChan:
				timer.Stop()
				// return function
				// don't not use break
				return
			}
			// break straight
			break
		}
	}
}

func (c *Cron) process(job FuncJob) {
	defer func() {
		if r := recover(); r != nil {
			golog.Error("job process panic recover")
			debug.PrintStack()
		}
	}()
	// run job
	job()
}

// Stop crontab
// if crontab already stopped no operator
func (c *Cron) Stop() {
	if !c.running {
		return
	}

	c.running = false
	// write 2 stop channel
	c.stopChan <- true
}
