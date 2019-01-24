// Package user update user
// Created by chenguolin 2018-11-16
package user

import (
	"gitlab.local.com/golang/httpserver/cmd/api/request"
	"gitlab.local.com/golang/httpserver/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateUserRequest 请求参数定义
type UpdateUserRequest struct {
	UID      int64  `json:"uid"       form:"uid"       binding:"required"`
	NewName  string `json:"new_name"  form:"new_name"  binding:"required"`
	NewPhone string `json:"new_phone" form:"new_phone" binding:"required"`
}

// UpdateUserRespone 响应参数定义
type UpdateUserRespone struct {
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	// 1. 参数校验
	req := new(UpdateUserRequest)
	err := c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
	if err != nil {
		errLog := "[update - UpdateUser] 参数无效, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 2. 更新用户信息
	userService := user.GetUserService()
	err = userService.UpdateUser(req.UID, req.NewName, req.NewPhone)
	if err != nil {
		errLog := "[update - UpdateUser] 更新用户信息出错, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 3. 响应response
	response := &UpdateUserRespone{}
	request.APISuccess(c, response)
}
