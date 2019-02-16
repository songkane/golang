// Package main HTTP Service router配置
// Created by chenguolin 2018-11-16
//
// 1. Get请求参数拼接在uri后面, 例如http://localhost:8080/user/info.json?uid=123456
// 2. Post请求参数使用key=value方式，例如"uid=123456&name=chenguolin"
// 3. content-type默认为 application/x-www-form-urlencoded
package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-http/cmd/api/controller"
	"gitlab.local.com/golang/go-http/cmd/api/controller/user"
)

// SetupRoute 设置路由
func SetupRoute(r *gin.Engine) {
	// 1. 设置业务相关的中间件，可以是全局的也可以是某一组APi特有
	r.Use(controller.SetRequestID())
	r.Use(controller.SetCommonParams())
	r.Use(controller.CheckSignature())
	r.Use(controller.CheckAuthorization())

	// 2. router设置
	// 使用分组设置router
	userGroup := r.RouterGroup.Group("/user")
	userGroup.GET("/select.json", user.SelectUserInfo)
	userGroup.POST("/create.json", user.CreateUser)
	userGroup.POST("/update.json", user.UpdateUser)
	userGroup.POST("/delete.json", user.DeleteUser)
}
