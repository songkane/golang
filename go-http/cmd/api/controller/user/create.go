// Package user create user api
// Created by chenguolin 2018-11-16
package user

import (
	"gitlab.local.com/golang/go-http/cmd/api/request"
	"gitlab.local.com/golang/go-http/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUserRequest 请求参数
type CreateUserRequest struct {
	UID   int64  `json:"uid"   form:"uid"   binding:"required"`
	Name  string `json:"name"  form:"name"  binding:"required"`
	Phone string `json:"phone" form:"phone" binding:"required"`
}

// CreateUserResponse 响应参数
type CreateUserResponse struct {
}

// CreateUser 创建一个新用户
func CreateUser(c *gin.Context) {
	// 1. 参数校验
	req := new(CreateUserRequest)
	err := c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
	if err != nil {
		errLog := "[create - CreateUser] 参数无效, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 2. 创建用户
	userService := user.GetUserService()
	args := &user.AddUserArgs{
		UID:   req.UID,
		Name:  req.Name,
		Phone: req.Phone,
	}

	err = userService.AddUser(args)
	if err != nil {
		errLog := "[create - CreateUser] 创建新用户失败, error: " + err.Error()
		request.APIError(c, errLog)
		return
	}

	// 3. 响应response
	response := &CreateUserResponse{}
	request.APISuccess(c, response)
}
