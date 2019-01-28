// Package signature http参数签名算法
// Created by chenguolin 2018-11-17
package signature

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// HTTPSigSdkAppSecret 签名秘钥
	HTTPSigSdkAppSecret = "abcdefghijklnmop"
	// HTTPSigSdkAddKey 签名盐
	HTTPSigSdkAddKey = "abcdefghijklnmop"
	// Sig 字段
	Sig = "sig"
	// SigTime 字段
	SigTime = "sigTime"
	// SigVersion 字段
	SigVersion = "sigVersion"
	// SigFinalString 字段
	SigFinalString = "sigFinalString"
)

// GetSignature 计算签名
func GetSignature(c *gin.Context) string {
	path := strings.TrimLeft(c.Request.URL.Path, "/")
	params := sortFormValues(c.Request.Form)
	sigTime := c.Request.Form.Get(SigTime)
	return generateSignature(path, params, sigTime)
}

func sortFormValues(formValues url.Values) []string {
	params := make([]string, 0, 10)
	for k, v := range formValues {
		if Sig == k || SigTime == k || SigVersion == k || SigFinalString == k {
			continue
		}
		params = append(params, v[0])
	}
	sort.Strings(params)

	return params
}

func generateSignature(path string, paramsArray []string, sigTime string) string {
	str := path + strings.Join(paramsArray, "") + HTTPSigSdkAppSecret + sigTime
	str += HTTPSigSdkAddKey
	temp := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	sig := bytes.Buffer{}
	var pos int
	for i := 0; i < 16; i++ {
		pos = i * 2
		sig.WriteByte(temp[pos+1])
		sig.WriteByte(temp[pos])
	}
	return sig.String()
}
