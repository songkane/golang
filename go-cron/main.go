// Package main go crontab
// Created by chenguolin 2018-11-17
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.local.com/golang/go-common/logger"
	"gitlab.local.com/golang/go-cron/cron"
	"gitlab.local.com/golang/go-cron/handle"
	golog "gitlab.local.com/golang/go-log"
)

const (
	// AppName application name
	AppName = "go-cron"
	// TimePattern default time pattern
	TimePattern = "20060102-15"
)

// Flags cmd args
type Flags struct {
	logDir string
}

// parseFlags parse cmd args
func parseFlags() *Flags {
	flags := new(Flags)
	flag.StringVar(&flags.logDir, "log_dir", "./logs", "the logger dir")
	flag.Parse()
	return flags
}

// main entry
func main() {
	// 1. parse flags
	flags := parseFlags()

	// 2. init logger
	log, err := logger.NewGolog(flags.logDir, AppName, TimePattern)
	if err != nil {
		panic(fmt.Sprintf("logger.NewGolog error:%s", err.Error()))
	}
	golog.SetLogger(log)

	// 3. start cron
	c := cron.NewCron()

	// TODO add handle
	c.AddHandle(cron.NewScheduler(cron.WithSecond(5), time.Now()), handle.GetCrawlAddressTxsHandle())

	// 4. run cron
	c.Start()
	fmt.Println("Start cron handler ...")

	// 5. shutdown
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stopSignalChan
	if sig != nil {
		fmt.Println("got system signal:" + sig.String() + ", going to shutdown crontab")
		// Stop the scheduler (does not stop any jobs already running).
		c.Stop()
	}
	fmt.Println("Stop cron handler ~")
}
