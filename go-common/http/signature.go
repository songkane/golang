// Package http signature
// Created by chenguolin 2018-11-17
package http

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// httpSignatureSecret 秘钥
	httpSignatureSecret = "abcdefghijklnmop"
	// httpSignatureSalt 盐
	httpSignatureSalt = "abcdefghijklnmop"
	// sigField 字段
	sigField = "sig"
)

// GenSignature get http signature
// 1. get url path
// 2. sort request form params
// 3. combine params + secret + sigTime + salt
// 4. calculator md5
// 5. shuffle md5 byte
func GenSignature(c *gin.Context) string {
	// 1. get url path
	// http://localhost:8080/user/select
	path := strings.TrimLeft(c.Request.URL.Path, "/")

	// 2. sort form values
	// content-type muse be application/x-www-form-urlencoded
	params := make([]string, 0, 10)
	form := c.Request.Form
	for k, v := range form {
		// filter sig field
		if k == sigField {
			continue
		}
		params = append(params, v[0])
	}
	sort.Strings(params)

	// 3. calculator signature
	// combine
	str := path + strings.Join(params, "") + httpSignatureSecret + httpSignatureSalt
	// md5
	temp := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	// shuffle byte
	sig := bytes.Buffer{}
	var pos int
	for i := 0; i < 16; i++ {
		pos = i * 2
		sig.WriteByte(temp[pos+1])
		sig.WriteByte(temp[pos])
	}

	return sig.String()
}
