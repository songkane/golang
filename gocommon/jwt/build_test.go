// Package jwt 单元测试
// Modified by chenguolin 2018-12-06
package jwt

import (
	"testing"
)

func TestBuild(t *testing.T) {
	// case 1
	uid := int64(0)
	token, err := Build(uid)
	if err != nil {
		t.Fatal("jwt TestBuild case 1 err != nil")
	}
	if token == "" {
		t.Fatal("jwt TestBuild case 1 token == \"\"")
	}

	// case 2
	uid = int64(123456)
	token, err = Build(uid)
	if err != nil {
		t.Fatal("jwt TestBuild case 2 err != nil")
	}
	if token == "" {
		t.Fatal("jwt TestBuild case 2 token == \"\"")
	}
}
