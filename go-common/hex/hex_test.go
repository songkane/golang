// Package hex 单元测试
// Created by chenguolin 2018-11-17
package hex

import (
	"testing"
)

func TestHex2BigIntStr(t *testing.T) {
	// case 1
	hex := ""
	bigIntValue := Hex2BigIntStr(hex)
	if bigIntValue != "" {
		t.Fatal("TestHex2BigIntStr case 1 get bigIntValue != \"\"")
	}

	// case 2
	hex = "0x628890"
	bigIntValue = Hex2BigIntStr(hex)
	if bigIntValue != "6457488" {
		t.Fatal("TestHex2BigIntStr case 2 get bigIntValue != \"6457488\"")
	}
}

func TestHexDecode(t *testing.T) {
	// case 1
	hex := ""
	decodeValue := Decode(hex)
	if decodeValue != "" {
		t.Fatal("TestHexDecode case 1 decodeValue != \"\"")
	}

	// case 2
	hex = "0x326667623334"
	decodeValue = Decode(hex)
	if decodeValue != "2fgb34" {
		t.Fatal("TestHexDecode case 2 decodeValue != \"2fgb34\"")
	}
}
