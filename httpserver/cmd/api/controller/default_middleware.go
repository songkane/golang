// Package controller default middleware
// Created by chenguolin 2018-11-16
package controller

import (
	"github.com/gin-gonic/gin"
)

// DefaultMiddleware 默认中间件
func DefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
