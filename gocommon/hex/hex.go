// Package hex hex基础库
// Created by chenguolin 2018-11-17
package hex

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Hex2BigIntStr 16进制转成bigInt
func Hex2BigIntStr(hex string) string {
	if hex == "" {
		return string("")
	}

	// 1. 16进制转成bigint
	bigInt, err := hexutil.DecodeBig(hex)
	if err != nil {
		return string("")
	}

	return bigInt.String()
}

// Decode 16进制解码
func Decode(hex string) string {
	if hex == "" {
		return string("")
	}

	bytes, err := hexutil.Decode(hex)
	if err != nil {
		return string("")
	}

	return string(bytes)
}
