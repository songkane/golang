// Package main internal模块main函数入口
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
	"gitlab.local.com/golang/go-http/version"
	golog "gitlab.local.com/golang/go-log"
)

const (
	// AccessLogName HTTP access log file name
	AccessLogName = "access.log"
	// AppName application name
	AppName = "go-http-internal"
	// TimePattern default time pattern
	TimePattern = "20060102-15"
)

func main() {
	// 1. parse cmd args
	flags := parseFlags()
	// 是否只打印版本信息
	if flags.logDir == "" || flags.confFile == "" || flags.version {
		printVersion()
		return
	}
	// 2. init config file
	conf := initConf(flags.confFile)
	// 3. init logger
	initLogger(flags.logDir)
	// 4. init application context
	instance.AppInit(flags.confFile)
	// 5. start HTTP Server
	startHTTPServer(conf.Deploy.InternalAddr, flags.logDir)
}

// printVersion print version
func printVersion() {
	fmt.Println("")
	fmt.Println("Usage: ./go-http-internal -log_dir=./logs -conf=../conf/config.toml")
	fmt.Println("")
	fmt.Printf("Version: %s\nbuilt at: %s\ncommit: %s\n", version.Version, version.BuildDate, version.BuildCommit)
	fmt.Println("")
}

// Flags cmd args
type Flags struct {
	logDir   string
	confFile string
	version  bool
}

// parseFlags 启动命令参数解析
func parseFlags() *Flags {
	flags := new(Flags)
	flag.StringVar(&flags.logDir, "log_dir", "", "log dir")
	flag.StringVar(&flags.confFile, "conf", "", "config file")
	flag.BoolVar(&flags.version, "version", false, "version")
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
			return
		}
	}()
}
