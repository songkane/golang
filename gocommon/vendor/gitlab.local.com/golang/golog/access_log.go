// Package golog HTTP access log
// Created by chenguolin 2018-12-27
package golog

import (
	"time"

	"github.com/gin-gonic/gin"
)

// AccessLogFunc 记录HTTP请求日志
func AccessLogFunc(log *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// Process request
		if c != nil {
			c.Next()
		}
		// 记录access log
		log.Info("HTTP Access Log",
			Object("reqUrl", c.Request.URL),
			Object("reqForm", c.Request.Form),
			Object("reqHeader", c.Request.Header),
			Object("reqBody", c.Request.Body),
			String("reqUri", c.Request.RequestURI),
			Int64("reqContentLength", c.Request.ContentLength),
			String("reqHost", c.Request.Host),
			String("reqMethod", c.Request.Method),
			String("reqProto", c.Request.Proto),
			String("reqRemoteAddr", c.Request.RemoteAddr),
			Duration("reqLatency", time.Since(start)),
			Object("resHeader", c.Writer.Header()),
			Int("resStatus", c.Writer.Status()),
			Int("resSize", c.Writer.Size()),
		)
	}
}
