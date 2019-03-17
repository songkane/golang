// Package prometheus golang prometheus
// Created by chenguolin 2019-03-16
package prometheus

import (
	"fmt"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// init function default call by go
func init() {
	// start HTTP Server
	// open /metrics api use for prometheus pull metrics
	engine := gin.New()
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// recovery
	engine.Use(gin.Recovery())
	go func() {
		httpServer := &http.Server{
			// default use 8888 port
			Addr:    ":8888",
			Handler: engine,
		}
		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Start Prometheus HTTP Server failed:", err)
			os.Exit(0)
		}
	}()
}
