// Package user select user api
// Created by chenguolin 2018-11-16
package user

import (
	"gitlab.local.com/golang/httpserver/cmd/api/request"
	"gitlab.local.com/golang/httpserver/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SelectUserInfoRequest 请求参数定义
type SelectUserInfoRequest struct {
	UID int64 `json:"uid" form:"uid" binding:"required"`
}

// SelectUserInfoResponse 响应结果定义
type SelectUserInfoResponse struct {
	*user.InfoResult
}

// SelectUserInfo 获取用户信息
func SelectUserInfo(c *gin.Context) {
	// 1. 参数校验
	req := new(SelectUserInfoRequest)
	err := c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
	if err != nil {
		errLog := "[select - SelectUserInfo] 参数无效, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 2. 查询用户信息
	userService := user.GetUserService()
	userInfoResult, err := userService.SelectUser(req.UID)
	if err != nil {
		errLog := "[select - SelectUserInfo] 查询用户信息出错, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 3. 构建response
	respones := &SelectUserInfoResponse{
		userInfoResult,
	}
	request.APISuccess(c, respones)
}
