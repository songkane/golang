// Package controller default middleware
// Created by chenguolin 2018-11-16
package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-common/uuid"
	"gitlab.local.com/golang/go-http/cmd/api/request"
)

// DefaultMiddleware 业务通用中间件
func DefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set request id
		reqID := uuid.NewUlid()
		c.Set(request.RequestIDKey, reqID)

		// TODO 业务逻辑 一般会包括以下几点

		// 1. 特殊url过滤
		// 2. 公共参数校验
		// 3. 通用参数设置
		// 4. sig校验
		// 5. 权限校验

		// 下一个handler
		c.Next()
	}
}
