// Package executor 重试执行封装函数
// Created by chenguolin 2018-11-16
package executor

import (
	"time"

	golog "gitlab.local.com/golang/go-log"
)

// RunUntilSuccess 无限重试直到成功
// @name 执行步骤名称
// @f 执行函数
func RunUntilSuccess(name string, f func() (interface{}, error)) interface{} {
	for {
		// run function
		result, err := f()
		if err == nil {
			return result
		}

		golog.Error("RunUntilSuccess execute error",
			golog.String("name", name),
			golog.Object("error", err))
		// default sleep 3 second
		time.Sleep(3 * time.Second)
	}
}

// RunUntilSuccessNoRes 无限重试直到成功 没有返回结果
// @name 执行步骤名称
// @f 执行函数
func RunUntilSuccessNoRes(name string, f func() error) {
	for {
		// run function
		err := f()
		if err == nil {
			return
		}

		golog.Error("RunUntilSuccessNoRes execute error",
			golog.String("name", name),
			golog.Object("error", err))
		// default sleep 3 second
		time.Sleep(3 * time.Second)
	}
}
