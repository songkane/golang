// Package jwt jwt生成算法
// Created by chenguolin 2018-11-21
package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

// VerifyRes 验证响应
type VerifyRes struct {
	Code  int
	Error string
	Biz   string
	User  int64
}

// Verify 验证JWT
// token AccessToken
func Verify(token string) *VerifyRes {
	// token = [header payload signature]
	fields := strings.Split(token, ".")
	if len(fields) != 3 {
		return VerifyInvalidToken
	}

	header, _ := base64.StdEncoding.DecodeString(fields[0])
	payload, _ := base64.StdEncoding.DecodeString(fields[1])

	var headerStruct Header
	err := json.Unmarshal([]byte(header), &headerStruct)
	if err != nil {
		return VerifyInvalidTokenHeader
	}

	var payloadStruct Payload
	err = json.Unmarshal([]byte(payload), &payloadStruct)
	if err != nil {
		return VerifyInvalidTokenPayload
	}

	// 验证是否过期
	now := time.Now().Unix()
	if now > payloadStruct.Exp {
		// 过期
		return VerifyTokenExpire
	}

	// 发现加密方式与本服务定义的不符，所以该token肯定不是自己的
	if headerStruct.Alg != HeaderAlg {
		return VerifyInvalidHeaderAlg
	}

	// 加密处理
	signature := SignatureBuild(fields[0]+"."+fields[1], Secret)
	// 验证通过
	if signature == fields[2] {
		res := VerifyPass
		res.User = payloadStruct.User
		res.Biz = payloadStruct.Biz
		return res
	}

	// 验证不通过
	res := VerifyFailed
	res.User = payloadStruct.User
	res.Biz = payloadStruct.Biz
	return res
}
