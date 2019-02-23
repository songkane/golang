// Package sql hash
// Created by chenguolin 2019-02-23
package sql

import (
	"hash/crc32"
	"strconv"
)

// Crc32 input calculate crc
func Crc32(input int64) int64 {
	return int64(crc32.ChecksumIEEE([]byte(strconv.FormatInt(input, 10))))
}

// GetDbIndex get multi db index
func GetDbIndex(input int64, dbCount int) int {
	return int(Crc32(input)) % dbCount
}

// GetTableIndex get multi table index
func GetTableIndex(input int64, tableCount int) int {
	return int(Crc32(input)) % tableCount
}
