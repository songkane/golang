// Package request response
// Created by chenguolin 2018-11-17
package request

import (
	"github.com/gin-gonic/gin"

	"gitlab.local.com/golang/httpserver/common/ulid"
)

const (
	// RequestIDKey request_id
	RequestIDKey = "request_id"
)

// EmptyStruct 空结构体定义
type EmptyStruct struct {
}

// Meta response meta定义
type Meta struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	RequestID  string `json:"request_id"`
	RequestURI string `json:"request_uri"`
}

// Response 定义
type Response struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"response"`
}

// APIError api响应失败
func APIError(c *gin.Context, errLog string) {
	// TODO
}

// APISuccess api响应成功
func APISuccess(c *gin.Context, response interface{}) {
	// TODO
}

// GetRequestID 获取请求ID
func GetRequestID(c *gin.Context) string {
	reqIDValue, isExist := c.Get(RequestIDKey)
	var reqID string
	if isExist {
		reqID = reqIDValue.(string)
	} else {
		//生成RequestID
		reqID = ulid.NewUlid()
		c.Set(RequestIDKey, reqID)
	}
	return reqID
}
