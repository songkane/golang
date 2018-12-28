/*
Package golog API接口
Created by chenguolin 2018-12-26
*/
package golog

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
		panic(fmt.Sprintf("NewLogger error", err))
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
	goLog.Log.Debug(msg, fields...)
}

// Info log
func Info(msg string, fields ...zap.Field) {
	goLog.Log.Info(msg, fields...)
}

// Warn log
func Warn(msg string, fields ...zap.Field) {
	goLog.Log.Warn(msg, fields...)
}

// Error log
func Error(msg string, fields ...zap.Field) {
	goLog.Log.Error(msg, fields...)
}

// Panic log
func Panic(msg string, fields ...zap.Field) {
	goLog.Log.Panic(msg, fields...)
}

// Fatal log
func Fatal(msg string, fields ...zap.Field) {
	goLog.Log.Fatal(msg, fields...)
}
