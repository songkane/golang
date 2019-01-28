// Package jwt jwt生成算法
// Created by chenguolin 2018-11-21
package jwt

import (
	"time"
)

// Build 根据用户id生成AccessToken
// uid 用户id
func Build(uid int64) (string, error) {
	// JWT加密算法采用SH256
	Token := build(HeaderAlg)
	Token.Payload.Exp = time.Now().Add(time.Hour * time.Duration(ExpHours)).Unix()
	Token.Payload.User = uid
	Token.Payload.Biz = Biz

	// 密钥
	jsn, err := Token.ToJSON(Secret)
	return jsn, err
}
