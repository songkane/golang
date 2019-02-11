// Package ticker ticker封装
// Created by chenguolin 2019-02-11
package ticker

import (
	"fmt"
	"runtime/debug"
	"time"
)

// Ticker run ticker
func Ticker(d time.Duration, f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	// 当主进程退出的时候 协程也会全部结束
	go func() {
		// run function
		f()

		// new ticker
		tk := time.NewTicker(d)
		defer tk.Stop()

		for {
			select {
			case <-tk.C:
				f()
			}
		}
	}()
}
