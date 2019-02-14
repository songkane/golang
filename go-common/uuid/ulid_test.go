// Package ulid 单元测试
// Created by chenguolin 2018-11-16
package uuid

import (
	"testing"
)

func TestNewUlid(t *testing.T) {
	ulid := NewUlid()
	if ulid == "" {
		t.Fatal("TestNewUlid NewUlid == \"\"")
	}
}
