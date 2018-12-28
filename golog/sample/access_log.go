/*
Package main HTTP service access log
Created by chenguolin 2018-12-27
*/
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golog"
)

// HelloServer the web server
func HelloServer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong successful ~",
	})
}

func main() {
	// new rotate writer
	fileName := os.Getenv("GOPATH") + "/src/gitlab.local.com/golog/sample/access.log"
	// 按小时切割文件
	writer, err := golog.NewRotateWriter(fileName, "20060102-15")
	if err != nil {
		fmt.Println("golog NewRotateWriter error", err)
	}

	// new logger
	log, err := golog.NewLogger(golog.WithOutput(writer), golog.WithJSONEncoder(), golog.WithInfoLevel())
	if err != nil {
		fmt.Printf("golog NewLogger error", err)
	}

	// start gin HTTP server
	r := gin.New()
	r.GET("/ping", golog.AccessLogFunc(log), HelloServer)
	// listen and serve on 0.0.0.0:8899
	r.Run(":8899")
}
