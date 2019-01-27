// Package time 时间基础库
// Created by chenguolin 2018-11-18
package time

import (
	"time"
)

// Sleep time
func Sleep(d time.Duration) {
	time.Sleep(d)
}

// GetTodayZeroTimestamp 获取今天零时时间戳
func GetTodayZeroTimestamp() int64 {
	curTime := time.Now()
	newTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(),
		0, 0, 0, 0, curTime.Location())
	return newTime.UnixNano() / int64(time.Second)
}

// GetCurrentTime 获取现在时间
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Timestamp2Date 时间戳转成日期
func Timestamp2Date(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}

// GetDayBeginTime 获取一天开始的时间
func GetDayBeginTime(curTime time.Time) time.Time {
	return time.Date(curTime.Year(), curTime.Month(), curTime.Day(),
		0, 0, 0, 0, curTime.Location())
}

// GetDayEndTime 获取一天结束的时间
func GetDayEndTime(curTime time.Time) time.Time {
	dayBeginTime := GetDayBeginTime(curTime)
	return dayBeginTime.Add(24 * time.Hour)
}
