// Package main HTTP Service router配置
// Created by chenguolin 2018-11-16
//
// 1. Get请求参数拼接在uri后面, 例如http://localhost:8080/user/info.json?uid=123456
// 2. Post请求参数使用key=value方式，例如"uid=123456&name=chenguolin"
// 3. content-type默认为 application/x-www-form-urlencoded
package main

import (
	"gitlab.local.com/golang/go-http/cmd/api/controller"
	"gitlab.local.com/golang/go-http/cmd/api/controller/user"

	"github.com/gin-gonic/gin"
)

// SetupRoute 设置路由
func SetupRoute(r *gin.Engine) {
	// 1. 设置业务全局的中间件
	// 这里的middleware 执行顺序在全局的recovery和accesslog middleware
	r.Use(controller.DefaultMiddleware())

	// 2. router设置
	// 使用分组设置router
	userGroup := r.RouterGroup.Group("/user")
	userGroup.GET("/select.json", user.SelectUserInfo)
	userGroup.POST("/create.json", user.CreateUser)
	userGroup.POST("/update.json", user.UpdateUser)
	userGroup.POST("/delete.json", user.DeleteUser)
}
