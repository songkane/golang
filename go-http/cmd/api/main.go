// Package main api模块main函数入口
// Created by chenguolin 2018-11-18
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	t "time"

	"github.com/gin-gonic/gin"

	"gitlab.local.com/golang/go-common/logger"
	"gitlab.local.com/golang/go-common/time"
	"gitlab.local.com/golang/go-http/config"
	"gitlab.local.com/golang/go-http/instance"
	golog "gitlab.local.com/golang/go-log"
)

const (
	// AccessLogName HTTP access log file name
	AccessLogName = "access.log"
	// AppName application name
	AppName = "go-http-api"
	// TimePattern default time pattern
	TimePattern = "20060102-15"
)

func main() {
	// 1. parse cmd args
	flags := parseFlags()
	// 2. init config file
	conf := initConf(flags.confFile)
	// 3. init logger
	initLogger(flags.logDir)
	// 4. init application context
	instance.AppInit(flags.confFile)
	// 5. init HTTP Server
	initHTTPServer(conf.Deploy.APIAddr, flags.logDir)
	// 6. start devops monitor server
	// TODO (@cgl)
	// devops.StartDevopsMonitorServer(conf.Deploy.DevopsAddr)
	// 7. block until HTTP Server shutdown
	blockUntilShutdown()
}

// Flags cmd args
type Flags struct {
	logDir   string
	confFile string
}

// parseFlags 启动命令参数解析
func parseFlags() *Flags {
	flags := new(Flags)
	flag.StringVar(&flags.logDir, "log_dir", "./logs", "log dir")
	flag.StringVar(&flags.confFile, "conf", "./conf/config.toml", "config file")
	flag.Parse()
	return flags
}

// initConf 配置文件初始化
func initConf(filePath string) *config.Config {
	apiConf := config.GetConfig(filePath)
	return apiConf
}

// initLogger 配置logger
func initLogger(logDir string) {
	log, err := logger.NewGolog(logDir, AppName, TimePattern)
	if err != nil {
		panic(fmt.Sprintf("initLogger logger.NewGolog error:%s", err.Error()))
	}
	golog.SetLogger(log)
}

// HTTPServer 变量
var (
	httpServer *http.Server
)

// initHTTPServer 初始化HTTPServer
func initHTTPServer(listenAddr string, accessLogDir string) {
	engine := gin.New()
	// gin goroutine recover
	engine.Use(gin.Recovery())
	// set access log
	log, err := logger.NewGolog(accessLogDir, AccessLogName, TimePattern)
	if err != nil {
		panic(fmt.Sprintf("initHTTPServer logger.NewGolog error:%s", err.Error()))
	}
	engine.Use(golog.AccessLogFunc(log))
	// set Router
	SetupRoute(engine)
	go func() {
		httpServer = &http.Server{
			Addr:    listenAddr,
			Handler: engine,
		}
		fmt.Println(time.GetCurrentTime(), "listening and serving HTTP on "+listenAddr)
		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "http server start failed:", err)
			os.Exit(0)
		}
	}()
}

// blockUntilShutdown
func blockUntilShutdown() {
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stopSignalChan
	if sig != nil {
		fmt.Println(time.GetCurrentTime(), "got system signal:"+sig.String()+", going to shutdown.")
		// wait resource remove from nginx upstreams
		time.Sleep(t.Second * 10)

		// 关闭http服务
		err := shutdownHTTPServer()
		if err != nil {
			fmt.Fprintln(os.Stderr, "http server graceful shutdown failed", err)
		} else {
			fmt.Println(time.GetCurrentTime(), "http server graceful shutdown successfully.")
		}
	}
}

// shutdownHTTPServer
func shutdownHTTPServer() error {
	// Create a deadline to wait for server shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 5*t.Second)
	defer cancel()
	return httpServer.Shutdown(ctx)
}
