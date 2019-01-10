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
	Entries  []*entry  // Entries: all run Entry set
	Running  bool      // Running: crontab is running
	StopChan chan bool // StopChan: crontab stop channel
}

// entry define
type entry struct {
	scheduler *Scheduler // Schedule: job
	job       FuncJob    // Job: run function
	nextTime  time.Time  // NextTime: next run time
	preTime   time.Time  // PreTime: last run time
}

// FuncJob implement Job interface
type FuncJob func()

// New crontab
// default use time location
func New() *Cron {
	return &Cron{
		Entries:  make([]*entry, 0),
		Running:  false,
		StopChan: make(chan bool),
	}
}

// AddJob add new job 2 crontab
// expression crontab expression (https://en.wikipedia.org/wiki/Cron)
// job run function
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
	c.Entries = append(c.Entries, entry)
}

// Entrys []*entry helper struct
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
	if c.Running {
		return
	}
	if len(c.Entries) <= 0 {
		return
	}

	// set Running true
	c.Running = true
	c.run()
}

func (c *Cron) run() {
	for {
		// 1. sort all entry by NextTime
		sort.Sort(entrys(c.Entries))

		// 2. new timer
		timer := time.NewTimer(c.Entries[0].nextTime.Sub(time.Now()))

		// 3. start run job
		// process label
		for {
			select {
			case now := <-timer.C:
				// run all has same NextTime entry
				for _, e := range c.Entries {
					if e.nextTime.After(now) {
						break
					}
					// goroutine
					go c.process(e.job)
					e.preTime = now
					e.nextTime = e.scheduler.Next(now)
				}
			case <-c.StopChan:
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
	if !c.Running {
		return
	}

	c.Running = false
	// write 2 stop channel
	c.StopChan <- true
}
