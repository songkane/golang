// Package executor 重试执行封装函数
// Created by chenguolin 2018-11-16
package executor

import (
	"time"

	golog "gitlab.local.com/golang/go-log"
)

// TryUntilSuccess 无限重试直到成功
// stepName 执行步骤名称
// f 执行函数
func TryUntilSuccess(stepName string, f func() (interface{}, error)) interface{} {
	for {
		result, err := f()
		if err != nil {
			golog.Error("[executor - TryUntilSuccess] execute error",
				golog.String("step", stepName),
				golog.Object("error", err))
			time.Sleep(3 * time.Second)
		} else {
			return result
		}
	}
}

// TryIgnoreErr 重试几次 忽略错误
// stepName 执行步骤名称
// f 执行函数
func TryIgnoreErr(stepName string, f func() error, retryTimes int) {
	for i := 0; i < retryTimes; i++ {
		err := f()
		if err != nil {
			golog.Error("[executor - TryIgnoreErr] execute error",
				golog.String("step", stepName),
				golog.Object("error", err))

			if i < retryTimes-1 {
				time.Sleep(2 * time.Second)
			}
		} else {
			return
		}
	}
}
