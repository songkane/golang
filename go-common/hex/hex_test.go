// Package hex 单元测试
// Created by chenguolin 2018-11-17
package hex

import (
	"testing"
)

func TestHex2BigIntStr(t *testing.T) {
	// case 1
	hex := ""
	bigIntValue, err := Hex2BigInt(hex)
	if err == nil {
		t.Fatal("TestHex2BigInt case 1 err == nil")
	}
	if bigIntValue != nil {
		t.Fatal("TestHex2BigInt case 1 get bigIntValue != nil")
	}

	// case 2
	hex = "0x628890"
	bigIntValue, err = Hex2BigInt(hex)
	if err != nil {
		t.Fatal("TestHex2BigInt case 2 err != nil")
	}
	if bigIntValue.String() != "6457488" {
		t.Fatal("TestHex2BigInt case 2 get bigIntValue.String() != \"6457488\"")
	}
}

func TestHexDecode(t *testing.T) {
	// case 1
	hex := ""
	decodeValue := DecodeHex(hex)
	if decodeValue != "" {
		t.Fatal("TestHexDecode case 1 decodeValue != \"\"")
	}

	// case 2
	hex = "0x326667623334"
	decodeValue = DecodeHex(hex)
	if decodeValue != "2fgb34" {
		t.Fatal("TestHexDecode case 2 decodeValue != \"2fgb34\"")
	}
}
