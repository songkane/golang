// Package golog unit test
// Created by chenguolin 2018-12-25
package golog

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	logPath   = os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/golog/logs"
	accessLog = logPath + "/" + "access.log"
)

func testInit() *http.Server {
	// mkdir log dir
	_ = os.MkdirAll(logPath, 0755)

	// new rotate writer
	// 按小时切割文件
	writer, err := NewRotateWriter(accessLog, "20060102-15")
	if err != nil {
		panic(fmt.Sprintf("TestAccessLogFunc golog NewRotateWriter error:%s", err.Error()))
	}

	// new logger
	log, err := NewLogger(WithOutput(writer), WithJSONEncoder(), WithInfoLevel())
	if err != nil {
		panic(fmt.Sprintf("TestAccessLogFunc golog NewLogger error:%s", err.Error()))
	}

	// start gin HTTP server
	r := gin.New()
	r.GET("/ping", AccessLogFunc(log), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping successful ~",
		})
	})
	// listen and serve on 0.0.0.0:8899
	srv := &http.Server{
		Addr:    ":8899",
		Handler: r,
	}

	return srv
}

func testDestroy(srv *http.Server) {
	// shutdown http server
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// srv.Shutdown(ctx)

	// remove dir
	_ = os.RemoveAll(logPath)
}

func TestAccessLogFunc(t *testing.T) {
	// test init
	srv := testInit()
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			t.Fatal("TestAccessLogFunc http server ListenAndServe error", err)
		}
	}()

	// http get request
	resp, err := http.Get("http://localhost:8899/ping")
	if err != nil {
		t.Fatal("TestAccessLogFunc http get error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestAccessLogFunc ioutil readall error", err)
	}
	if len(body) <= 0 {
		t.Fatal("TestAccessLogFunc ioutil readall get body empty")
	}

	// check log file content
	f, err := os.Open(accessLog)
	if err != nil {
		t.Fatal("TestAccessLogFunc os open access log file error", err)
	}
	defer f.Close()

	content := make([]byte, 4096)
	size, err := f.Read(content)
	if err != nil {
		t.Fatal("TestAccessLogFunc f.Read error", err)
	}
	// {"Level":"INFO","TimeStamp":"2019-01-03 22:01:40.639815265","Name":"ZapJSONLogger","Message":"HTTP Access Log","reqUrl":"/ping","reqForm":null,"reqHeader":{"Accept-Encoding":["gzip"],"User-Agent":["Go-http-client/1.1"]},"reqBody":{},"reqUri":"/ping","reqContentLength":0,"reqHost":"localhost:8899","reqMethod":"GET","reqProto":"HTTP/1.1","reqRemoteAddr":"127.0.0.1:60987","reqLatency":"102.231µs","resHeader":{"Content-Type":["application/json; charset=utf-8"]},"resStatus":200,"resSize":31}
	fmt.Println(string(content))
	if size <= 0 {
		t.Fatal("TestAccessLogFunc f.Read size <= 0")
	}

	// test destroy
	testDestroy(srv)
}
