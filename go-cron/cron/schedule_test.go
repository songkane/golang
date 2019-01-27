// Package cron golang crontab
// crontab scheduler
// Created by chenguolin 2018-01-04
package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestNewScheduler(t *testing.T) {
	sche := NewScheduler(WithSecond(5), time.Now())
	if sche == nil {
		t.Fatal("TestNewScheduler sche == nil")
	}
}

func TestWithSecond(t *testing.T) {
	sche := NewScheduler(WithSecond(5), time.Now())
	if sche == nil {
		t.Fatal("TestWithSecond sche == nil")
	}

	if sche.Next(sche.startTime).Sub(sche.startTime) != time.Duration(5)*time.Second {
		t.Fatal("TestWithSecond test failed")
	}
}

func TestWithMinute(t *testing.T) {
	sche := NewScheduler(WithMinute(1), time.Now())
	if sche == nil {
		t.Fatal("TestWithMinute sche == nil")
	}

	if sche.Next(sche.startTime).Sub(sche.startTime) != time.Duration(1)*time.Minute {
		t.Fatal("TestWithMinute test failed")
	}
}

func TestWithHour(t *testing.T) {
	sche := NewScheduler(WithHour(2), time.Now())
	if sche == nil {
		t.Fatal("TestWithHour sche == nil")
	}

	if sche.Next(sche.startTime).Sub(sche.startTime) != time.Duration(2)*time.Hour {
		t.Fatal("TestWithHour test failed")
	}
}

func TestWithDay(t *testing.T) {
	sche := NewScheduler(WithDay(3), time.Now())
	if sche == nil {
		t.Fatal("TestWithDay sche == nil")
	}

	if sche.Next(sche.startTime).Sub(sche.startTime) != time.Duration(3*24)*time.Hour {
		t.Fatal("TestWithDay test failed")
	}
}

func TestScheduler_Next(t *testing.T) {
	sche := NewScheduler(WithSecond(5), time.Now())
	if sche == nil {
		t.Fatal("TestScheduler_Next sche == nil")
	}
	fmt.Println(sche.Next(sche.startTime))
}
