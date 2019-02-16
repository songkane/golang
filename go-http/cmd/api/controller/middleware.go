// Package controller default middleware
// Created by chenguolin 2018-11-16
package controller

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.local.com/golang/go-common/http"
	"gitlab.local.com/golang/go-common/uuid"
	"gitlab.local.com/golang/go-http/cmd/api/request"
)

// SetRequestID set request id
func SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set request id
		reqID := uuid.NewUlid()
		c.Set(request.RequestIDKey, reqID)
		c.Next()
	}
}

// SetCommonParams set common params
func SetCommonParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 业务逻辑 一般会包括以下几点
		// 1. 特殊url过滤
		// 2. 公共参数校验
		// 3. 通用参数设置

		// 下一个handler
		c.Next()
	}
}

// CheckSignature check http signature
// Signature header format:
//    `Signature: sigTime={sigTime}&signature={signature}`
// variables:
// - sigTime = unix timestamp example 1547890889
// - signature = signature string
func CheckSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get signature header
		sigHeader := c.GetHeader("Signature")
		fields := strings.Split(sigHeader, "&")
		// check fields
		if fields == nil || len(fields) <= 2 {
			errLog := string("Signature header invalid ~")
			request.APIError(c, errLog)
			return
		}

		// 2. parse fields as query
		params, err := url.ParseQuery(sigHeader)
		if err != nil {
			errLog := string("Signature header parse error: ") + err.Error()
			request.APIError(c, errLog)
			return
		}

		// 3. check sigTime
		sigTimeStr := params.Get("sigTime")
		sigTime, err := strconv.ParseInt(sigTimeStr, 10, 64)
		if err != nil {
			errLog := string("Signature sigTimeStr ParseInt error: ") + err.Error()
			request.APIError(c, errLog)
			return
		}
		if sigTime <= time.Now().Unix() {
			errLog := string("Signature sigTime expired ~")
			request.APIError(c, errLog)
			return
		}
		// set sigTime
		c.Set("sigTime", sigTimeStr)

		// 4. check signature
		signature := params.Get("signature")
		if signature != http.GenSignature(c) {
			errLog := string("Signature check failed ~")
			request.APIError(c, errLog)
			return
		}

		// 下一个handler
		c.Next()
	}
}

// CheckAuthorization check http authorization
// Authorization header format:
//    `Authorization: pubk={pubk}&sigTime={sigTime}&signature={signature}`
// variables:
// - pubk = public key
// - sigTime = unix timestamp example 1547890889
// - signature = signature string
func CheckAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get signature header
		authHeader := c.GetHeader("Authorization")
		fields := strings.Split(authHeader, "&")
		// check fields
		if fields == nil || len(fields) <= 2 {
			errLog := string("Authorization header invalid ~")
			request.APIError(c, errLog)
			return
		}

		// 2. parse fields as query
		params, err := url.ParseQuery(authHeader)
		if err != nil {
			errLog := string("Authorization header parse error: ") + err.Error()
			request.APIError(c, errLog)
			return
		}

		// 3. check sigTime
		sigTimeStr := params.Get("sigTime")
		sigTime, err := strconv.ParseInt(sigTimeStr, 10, 64)
		if err != nil {
			errLog := string("Authorization sigTimeStr ParseInt error: ") + err.Error()
			request.APIError(c, errLog)
			return
		}
		if sigTime <= time.Now().Unix() {
			errLog := string("Authorization sigTime expired ~")
			request.APIError(c, errLog)
			return
		}

		// 4. get public key and signature
		pubk := params.Get("pubk")
		sig := params.Get("signature")

		// 4. check authrization
		c.Set("publicKey", pubk)
		c.Set("signature", sig)
		c.Set("sigTime", sigTimeStr)

		// 下一个handler
		c.Next()
	}
}
