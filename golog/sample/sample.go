package main

import (
	"os"

	"gitlab.local.com/golog"
)

func main() {
	// use 1
	// 默认level为Info
	// 默认使用JSONEncoder
	// 默认输出到Stdout
	golog.Info("use 1", golog.String("name", "chenguolin"))

	// use 2
	// level设置为Info
	// 使用Console Encodr
	// 输出到Stderr
	log, err := golog.NewLogger(golog.WithInfoLevel(), golog.WithConsoleEncoder(), golog.WithOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	golog.SetLogger(log)
	golog.Info("use 2", golog.String("name", "chenguolin"))
}
