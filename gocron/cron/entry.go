// Package cron entry
// Created by chenguolin 2018-01-04
package cron

import (
	"time"
)

// entry cron entry
type entry struct {
	scheduler *Scheduler // cron scheduler
	handle    handle     // cron handle
	nextTime  time.Time  // next run time
	preTime   time.Time  // last run time
}

// entrys []*entry helper type
type entrys []*entry

// Len return entrys length
func (es entrys) Len() int {
	return len(es)
}

// Swap entry[i] and entry[j]
func (es entrys) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

// Less check entry[i] next time before entry[j] next time
func (es entrys) Less(i, j int) bool {
	// if entry[i] next time is zero return false
	if es[i].nextTime.IsZero() {
		return false
	}

	if es[j].nextTime.IsZero() {
		return true
	}

	return es[i].nextTime.Before(es[j].nextTime)
}
