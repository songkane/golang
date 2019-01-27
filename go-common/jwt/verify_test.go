// Package jwt 单元测试
// Modified by chenguolin 2018-12-06
package jwt

import (
	"testing"
)

func TestVerify(t *testing.T) {
	uid := int64(123456)
	token, err := Build(uid)
	if err != nil {
		t.Fatal("TestVerify case 1 err != nil")
	}

	res := Verify(token)
	if res == nil {
		t.Fatal("TestVerify case 1 res == nil")
	}
	if res.Code != 0 {
		t.Fatal("TestVerify case 1 res.Code != 0")
	}
	if res.Error != "" {
		t.Fatal("TestVerify case 1 res.Error != \"\"")
	}
	if res.User != uid {
		t.Fatal("TestVerify case 1 res.User != uid")
	}
	if res.Biz != "httpserver" {
		t.Fatal("TestVerify case 1 res.User != \"httpserver\"")
	}
}
