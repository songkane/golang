// Package cron golang crontab
// Created by chenguolin 2018-01-04
package cron

import (
	"testing"
	"time"

	h "gitlab.local.com/golang/go-cron/handle"
)

func TestEntrys_Len(t *testing.T) {
	es := make(entrys, 0)

	sche := NewScheduler(WithSecond(2), time.Now())
	es1 := &entry{
		scheduler: sche,
		handle:    h.GetCrawlAddressTxsHandle(),
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(4), time.Now())
	es2 := &entry{
		scheduler: sche2,
		handle:    h.GetCrawlAddressTxsHandle(),
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
		handle:    h.GetCrawlAddressTxsHandle(),
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(8), time.Now())
	es2 := &entry{
		scheduler: sche2,
		handle:    h.GetCrawlAddressTxsHandle(),
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
		handle:    h.GetCrawlAddressTxsHandle(),
		nextTime:  sche.Next(time.Now()),
	}

	sche2 := NewScheduler(WithSecond(8), time.Now())
	es2 := &entry{
		scheduler: sche2,
		handle:    h.GetCrawlAddressTxsHandle(),
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
