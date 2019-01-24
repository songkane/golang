// Package bizerror biz error define
// Created by chenguolin 2018-11-17
package bizerror

import "errors"

var (
	// ErrInvalidArguments 通用error定义
	ErrInvalidArguments = errors.New("Invalid Arguments Error")
	// ErrMysqlRecordNotFound 数据库业务error定义
	ErrMysqlRecordNotFound = errors.New("Mysql Record Not Found Error")
	// ErrMysqlAffectedNoRows 数据库操作影响行数为0
	ErrMysqlAffectedNoRows = errors.New("Mysql Affected No Rows")
)
