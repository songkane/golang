// Package cron golang crontab
// Created by chenguolin 2018-01-04
package cron

import (
	"testing"
	"time"

	h "gitlab.local.com/golang/go-cron/handle"
)

func TestNewCron(t *testing.T) {
	cron := NewCron()
	if len(cron.entryList) > 0 {
		t.Fatal("TestNewCron len(cron.entries) > 0 ")
	}
	if cron.isRunning {
		t.Fatal("TestNewCron cron.isRunning == true ")
	}
	if len(cron.stopChan) > 0 {
		t.Fatal("TestNewCron len(cron.stopChan) > 0")
	}
}

func TestCron_AddHandle(t *testing.T) {
	cron := NewCron()
	sche := NewScheduler(WithSecond(2), time.Now())

	cron.AddHandle(sche, h.GetCrawlAddressTxsHandle())

	if len(cron.entryList) <= 0 {
		t.Fatal("TestCron_AddJob len(cron.entries) <= 0")
	}
}

func TestCron_Start(t *testing.T) {
	// case 1
	cron := NewCron()
	cron.Start()
	cron.Stop()

	// case 2
	sche := NewScheduler(WithSecond(2), time.Now())
	cron.AddHandle(sche, h.GetCrawlAddressTxsHandle())
	if len(cron.entryList) <= 0 {
		t.Fatal("TestCron_AddJob len(cron.entries) <= 0")
	}
	cron.Start()
	time.Sleep(time.Duration(6) * time.Second)
	cron.Stop()
}

func TestCron_Stop(t *testing.T) {
	// case 1
	cron := NewCron()
	cron.Stop()

	// case 2
	sche := NewScheduler(WithSecond(2), time.Now())
	cron.AddHandle(sche, h.GetCrawlAddressTxsHandle())
	if len(cron.entryList) <= 0 {
		t.Fatal("TestCron_AddJob len(cron.entries) <= 0")
	}
	cron.Start()
	time.Sleep(time.Duration(6) * time.Second)
	cron.Stop()
}
