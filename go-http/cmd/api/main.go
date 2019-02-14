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
	"gitlab.local.com/golang/go-healthcheck"
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
	// 5. start HTTP Server
	startHTTPServer(conf.Deploy.APIAddr, flags.logDir)
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

// startHTTPServer start HTTP Server
func startHTTPServer(listenAddr string, accessLogDir string) {
	// set access log
	log, err := logger.NewGolog(accessLogDir, AccessLogName, TimePattern)
	if err != nil {
		panic(fmt.Sprintf("initHTTPServer logger.NewGolog error:%s", err.Error()))
	}

	// new gin engine
	// set release mode
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// set gin global middleware handler
	// 1. recovery handler default write 2 os.stderr
	// 2. golog handler write http access log
	engine.Use(gin.Recovery())
	engine.Use(golog.AccessLogFunc(log))

	// register health and pprof
	healthcheck.RegisterHealthCheck(engine)

	// set Router
	SetupRoute(engine)

	// start HTTP Server
	httpServer := &http.Server{
		Addr:    listenAddr,
		Handler: engine,
	}
	fmt.Println(time.GetCurrentTime(), "Start HTTP server listening on "+listenAddr)
	go func() {
		err = httpServer.ListenAndServe()
		if err != nil {
			panic(fmt.Sprintf("HTTP Server start failed: %s", err.Error()))
		}
	}()

	// block until receive signal
	shutdown := make(chan struct{})
	registerSignal(shutdown)
	<-shutdown

	// shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*t.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
	fmt.Println(time.GetCurrentTime(), "Shutdown HTTP server successful ~")
}

// registerSignal register shutdown signal
func registerSignal(shutdown chan struct{}) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range c {
			close(shutdown)
			fmt.Println(time.GetCurrentTime(), "got system signal:"+sig.String()+", going to shutdown ...")
			break
		}
	}()
}
