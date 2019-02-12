// Package hex 16进制转换
// Created by chenguolin 2018-11-17
package hex

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Hex2BigInt hex 2 bigint
func Hex2BigInt(hex string) (*big.Int, error) {
	if hex == "" {
		return nil, errors.New("hex is empty")
	}

	// 1. 16进制转成bigint
	bigInt, err := hexutil.DecodeBig(hex)
	if err != nil {
		return nil, err
	}

	return bigInt, nil
}

// DecodeHex hex decode 2 string
func DecodeHex(hex string) string {
	if hex == "" {
		return string("")
	}

	bytes, err := hexutil.Decode(hex)
	if err != nil {
		return string("")
	}

	return string(bytes)
}
