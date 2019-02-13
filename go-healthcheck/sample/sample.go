// Package main sample
// Created by chenguolin 2019-02-13
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-healthcheck"
)

func main() {
	engine := gin.New()
	healthcheck.RegisterHealthCheck(engine)

	httpServer := &http.Server{
		Addr:    "0.0.0.0:6060",
		Handler: engine,
	}
	httpServer.ListenAndServe()
}
