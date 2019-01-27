// Package user delete user api
// Created by chenguolin 2018-11-16
package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gitlab.local.com/golang/go-http/cmd/api/request"
	"gitlab.local.com/golang/go-http/pkg/user"
)

// DeleteUserRequest 请求参数定义
type DeleteUserRequest struct {
	UID int64 `json:"uid"   form:"uid"   binding:"required"`
}

// DeleteUserResponse 响应结果定义
type DeleteUserResponse struct {
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// 1. 参数校验
	req := new(DeleteUserRequest)
	err := c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
	if err != nil {
		errLog := "[delete - DeleteUser] 参数无效, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 2. 删除用户
	userService := user.GetUserService()
	err = userService.DeleteUser(req.UID)
	if err != nil {
		errLog := "[delete - DeleteUser] 删除用户出错, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 3. 响应response
	response := &DeleteUserResponse{}
	request.APISuccess(c, response)
}
