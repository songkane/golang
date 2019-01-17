// Package cron golang crontab
// Created by chenguolin 2018-01-04
package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCron(t *testing.T) {
	cron := NewCron()
	if len(cron.entries) > 0 {
		t.Fatal("TestNewCron len(cron.entries) > 0 ")
	}
	if cron.running {
		t.Fatal("TestNewCron cron.running == true ")
	}
	if len(cron.stopChan) > 0 {
		t.Fatal("TestNewCron len(cron.stopChan) > 0")
	}
}

func TestCron_AddJob(t *testing.T) {
	cron := NewCron()
	sche := NewScheduler(WithSecond(2), time.Now())

	f := func() { fmt.Println("TestCron_AddJob") }
	cron.AddJob(sche, f)

	if len(cron.entries) <= 0 {
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
	f := func() { fmt.Println("TestCron_AddJob") }
	cron.AddJob(sche, f)
	if len(cron.entries) <= 0 {
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
	f := func() { fmt.Println("TestCron_AddJob") }
	cron.AddJob(sche, f)
	if len(cron.entries) <= 0 {
		t.Fatal("TestCron_AddJob len(cron.entries) <= 0")
	}
	cron.Start()
	time.Sleep(time.Duration(6) * time.Second)
	cron.Stop()
}

func TestEntrys_Len(t *testing.T) {
	es := make(entrys, 0)

	sche := NewScheduler(WithSecond(2), time.Now())
	es1 := &entry{
		scheduler: sche,
		job:       func() { fmt.Println("print es1") },
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(4), time.Now())
	es2 := &entry{
		scheduler: sche2,
		job:       func() { fmt.Println("print es2") },
		nextTime:  sche2.Next(time.Now()),
	}

	es = append(es, es1)
	es = append(es, es2)

	if len(es) != 2 {
		t.Fatal("TestEntrys_Len len(es) != 2")
	}
}

func TestEntrys_Swap(t *testing.T) {
	es := make(entrys, 0)

	sche := NewScheduler(WithSecond(5), time.Now())
	es1 := &entry{
		scheduler: sche,
		job:       func() { fmt.Println("print es1") },
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(8), time.Now())
	es2 := &entry{
		scheduler: sche2,
		job:       func() { fmt.Println("print es2") },
		nextTime:  sche2.Next(time.Now()),
	}

	es = append(es, es1)
	es = append(es, es2)

	es.Swap(0, 1)
	if es[0].scheduler != sche2 {
		t.Fatal("TestEntrys_Swap es[0].scheduler != sche2")
	}
	if es[1].scheduler != sche {
		t.Fatal("TestEntrys_Swap es[1].scheduler != sche1")
	}
}

func TestEntrys_Less(t *testing.T) {
	es := make(entrys, 0)

	sche := NewScheduler(WithSecond(5), time.Now())
	es1 := &entry{
		scheduler: sche,
		job:       func() { fmt.Println("print es1") },
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(8), time.Now())
	es2 := &entry{
		scheduler: sche2,
		job:       func() { fmt.Println("print es2") },
		nextTime:  sche2.Next(time.Now()),
	}

	es = append(es, es1)
	es = append(es, es2)

	isLess := es.Less(0, 1)
	if !isLess {
		t.Fatal("TestEntrys_Less es.Less(0, 1) return false")
	}

	isLess = es.Less(1, 0)
	if isLess {
		t.Fatal("TestEntrys_Less es.Less(1, 0) return true")
	}
}
