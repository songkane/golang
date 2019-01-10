// Package cron golang crontab
// crontab scheduler
// Created by chenguolin 2018-01-04
package cron

import (
	"time"
)

// SchedulerOption uint helper type
type SchedulerOption uint

// Scheduler define
type Scheduler struct {
	interval  time.Duration
	startTime time.Time
}

// WithSecond with second interval
func WithSecond(interval SchedulerOption) SchedulerOption {
	return interval
}

// WithMinute with minute interval
func WithMinute(interval SchedulerOption) SchedulerOption {
	return interval * 60
}

// WithHour with hour interval
func WithHour(interval SchedulerOption) SchedulerOption {
	return interval * 60 * 60
}

// WithDay with day interval
func WithDay(interval SchedulerOption) SchedulerOption {
	return interval * 60 * 60 * 24
}

// NewScheduler new scheduler
func NewScheduler(option SchedulerOption, startTime time.Time) *Scheduler {
	return &Scheduler{
		interval:  time.Duration(option),
		startTime: startTime,
	}
}

// Next schedule time
func (s *Scheduler) Next(t time.Time) time.Time {
	if s.startTime.After(t) {
		tmpT := t.Add(time.Duration(s.interval) * time.Second)
		if tmpT.Before(s.startTime) {
			return s.startTime
		}
		return tmpT
	}

	t = t.Add(time.Duration(s.interval) * time.Second)
	return t
}
