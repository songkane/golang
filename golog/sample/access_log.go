package main

import (
	"fmt"
	"io"
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
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd", err)
	}
	// set logger
	fileName := dir + "/sample/access.log"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Printf("open file error", err)
	}

	log, err := golog.NewLogger(golog.WithOutput(io.Writer(f)), golog.WithJSONEncoder(), golog.WithInfoLevel())
	if err != nil {
		fmt.Printf("golog NewLogger error", err)
	}

	// start gin HTTP server
	r := gin.New()
	r.GET("/ping", golog.AccessLogFunc(log), HelloServer)
	// listen and serve on 0.0.0.0:8899
	r.Run(":8899")
}
