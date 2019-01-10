// Package main go crontab
// Created by chenguolin 2018-11-17
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
	"gitlab.local.com/golang/gocommon/logger"
	"gitlab.local.com/golang/gocron/handler"
	"gitlab.local.com/golang/golog"
)

const (
	// AppName application name
	AppName = "gocron"
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
	c := cron.New()
	// all handlers
	// etc one by one
	c.AddFunc("0 30 * * * *", handler.CrawlAddressTxsHandler)

	// 4. run cron
	c.Run()

	// 5. shutdown
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stopSignalChan
	if sig != nil {
		fmt.Println("got system signal:" + sig.String() + ", going to shutdown crontab")
		// Stop the scheduler (does not stop any jobs already running).
		c.Stop()
	}
}
