// Package jwt 通用变量、常量定义
// Created by chenguolin 2018-11-21
package jwt

// 常量定义
const (
	// Secret 加密密钥
	Secret = "fjskfjls2io3242mn"
	// ExpHours 过期时间
	ExpHours = 12
	// 默认加密方式
	HeaderAlg = "SH256"
	// Biz 业务类型
	Biz = "httpserver"
)

// token验证错误定义
var (
	// VerifyPass token验证通过
	VerifyPass = &VerifyRes{
		Code:  0,
		Error: "",
	}
	// VerifyFailed token验证失败
	VerifyFailed = &VerifyRes{
		Code:  10000,
		Error: "verification failed",
	}
	// VerifyInvalidToken 无效token
	VerifyInvalidToken = &VerifyRes{
		Code:  20000,
		Error: "invalid token",
	}
	// VerifyInvalidTokenHeader token header无效
	VerifyInvalidTokenHeader = &VerifyRes{
		Code:  20001,
		Error: "invalid token header",
	}
	// VerifyInvalidTokenPayload token payload无效
	VerifyInvalidTokenPayload = &VerifyRes{
		Code:  20002,
		Error: "invalid token payload",
	}
	// VerifyTokenExpire token过期
	VerifyTokenExpire = &VerifyRes{
		Code:  20003,
		Error: "token expire",
	}
	// VerifyInvalidHeaderAlg token 无效加密方式
	VerifyInvalidHeaderAlg = &VerifyRes{
		Code:  20004,
		Error: "invalid token header algorithm",
	}
)
