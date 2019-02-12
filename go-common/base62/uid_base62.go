// Package base62 base62编码实现方式 给定一个整数编码成固定的6位字符
// Created by chenguolin 2018-08-22
package base62

import (
	"strings"
)

const (
	// addSaltValue 表示第一个盐值
	// md5("httpserver-base62-add") =》899c95e56f6c304416246279ba5675a9
	// 899c95e56f6c304416246279ba5675a9 =》3839396339356535366636633330343431363234363237396261353847349232
	// 取末尾10位
	addSaltValue = int64(3847349232)

	// xorSaltValue 第二位盐值
	// md5("httpserver-base62-xor") =》a86e0bf3281bee432572f0e269bb8a6c
	// a86e0bf3281bee432572f0e269bb8a6c =》6138366530626633323831626565343332353732663065323639627139412920
	// 取末尾10位
	xorSaltValue = int64(7139412920)

	// baseString 字符串集合 0~9 A~Z a~z 共计62个字符
	baseString = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	baseBytes = []byte(baseString)
	baseMap   = make(map[byte]int)
)

func init() {
	for i, v := range baseBytes {
		baseMap[v] = i
	}
}

// ID2Base62 把一个id编码成6位 base62字符
// 返回空串表示编码失败
func ID2Base62(id int64) string {
	if id <= 0 {
		return string("")
	}

	// 计算新的uid
	// 先加上盐 =》 A
	// 再异或另外一个盐 =》B
	// B =》base62 编码
	number := (id + addSaltValue) ^ xorSaltValue

	// 1. 10进制转62进制
	result := make([]byte, 0)
	for {
		if number <= 0 {
			break
		}
		mod := number % 62
		number = number / 62
		result = append(result, baseBytes[mod])
	}

	// reverse slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// 如果超过6位报错说明uid太大 返回空串
	// 如果不足6位需要前缀补0
	if len(result) > 6 {
		return string("")
	}

	// 计算需要补多少位0
	needAddZeroCnt := 6 - len(result)
	return strings.Repeat("0", needAddZeroCnt) + string(result)
}

// Base622ID 把一个6位base62编码字符串反解成id
// 返回-1表示解码失败
func Base622ID(base62 string) int64 {
	if base62 == "" {
		return -1
	}

	number := int64(0)
	pow := 0
	bytes := []byte(base62)

	for i := len(bytes) - 1; i >= 0; i-- {
		v, ok := baseMap[bytes[i]]
		if !ok {
			return -1
		}

		// 62进制转10进制
		tmpValue := int64(1)
		for j := 0; j < pow; j++ {
			tmpValue *= 62
		}

		// 加上mod的值
		number += tmpValue * int64(v)
		pow++
	}

	// 计算uid
	uid := (number ^ xorSaltValue) - addSaltValue
	return uid
}
