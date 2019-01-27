// Package golog API接口
// Created by chenguolin 2018-12-26
package go_log

import (
	"fmt"

	"go.uber.org/zap"
)

// 全局 logger
var goLog *Logger

func init() {
	// default logger
	lg, err := NewLogger()
	// 如果NewLogger出错直接panic 进程直接退出
	if err != nil {
		panic(fmt.Sprintf("NewLogger error: %s", err.Error()))
	}

	goLog = lg
}

// GetLogger get goLog
func GetLogger() *Logger {
	return goLog
}

// SetLogger set goLog
func SetLogger(log *Logger) {
	goLog = log
}

// Debug log
func Debug(msg string, fields ...zap.Field) {
	goLog.log.Debug(msg, fields...)
}

// Info log
func Info(msg string, fields ...zap.Field) {
	goLog.log.Info(msg, fields...)
}

// Warn log
func Warn(msg string, fields ...zap.Field) {
	goLog.log.Warn(msg, fields...)
}

// Error log
func Error(msg string, fields ...zap.Field) {
	goLog.log.Error(msg, fields...)
}

// Panic log
func Panic(msg string, fields ...zap.Field) {
	goLog.log.Panic(msg, fields...)
}

// Fatal log
func Fatal(msg string, fields ...zap.Field) {
	goLog.log.Fatal(msg, fields...)
}
