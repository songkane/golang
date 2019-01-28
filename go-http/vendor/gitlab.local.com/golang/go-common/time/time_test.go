// Package time 单元测试
// Created by chenguolin 2018-11-18
package time

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTodayZeroTimestamp(t *testing.T) {
	fmt.Println(GetTodayZeroTimestamp())
}

func TestGetCurrentTime(t *testing.T) {
	fmt.Println(GetCurrentTime())
}

func TestTimestamp2Date(t *testing.T) {
	fmt.Println(Timestamp2Date(1526357368))
}

func TestGetDayBeginTime(t *testing.T) {
	now := time.Date(2018, time.July, 10, 12, 0, 0, 0, time.FixedZone("UTC", 8*3600))
	begin := GetDayBeginTime(now)
	if begin.Unix() != 1531152000 {
		t.Fatal("GetDayBeginTime error", begin.Unix())
	}
}

func TestGetDayEndTime(t *testing.T) {
	now := time.Date(2018, time.July, 10, 12, 0, 0, 0, time.FixedZone("UTC", 8*3600))
	end := GetDayEndTime(now)
	if end.Unix() != 1531238400 {
		t.Fatal("GetDayEndTime error", end.Unix())
	}
}
