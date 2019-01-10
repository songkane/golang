// Package sql mysql 基础功能函数
// Created by chenguolin 2018-05-21
package sql

import (
	"strings"
)

// IsDuplicateKeyError 判断sql返回错误类型 "数据库主键重复"
func IsDuplicateKeyError(e error) bool {
	return e != nil && strings.Contains(e.Error(), "Duplicate entry")
}

// IsNotRowsError 判断sql返回错误类型 "没有结果"
func IsNotRowsError(e error) bool {
	return e != nil && strings.Contains(e.Error(), "sql: no rows")
}
