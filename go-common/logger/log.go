// Package logger 封装traceId的log打印
// Created by chenguolin 2018-11-17
package logger

import (
	"errors"
	"fmt"

	"gitlab.local.com/golang/go-common/trace"
	"gitlab.local.com/golang/go-log"
	"go.uber.org/zap"
)

// NewGolog new golog logger
func NewGolog(dir, name, timePattern string) (*golog.Logger, error) {
	if dir == "" || name == "" || timePattern == "" {
		return nil, errors.New("Invalid Arguments")
	}

	fileName := fmt.Sprintf("%s/%s.logger", dir, name)
	rotateWriter, err := golog.NewRotateWriter(fileName, timePattern)
	if err != nil {
		return nil, err
	}

	logger, err := golog.NewLogger(golog.WithOutput(rotateWriter), golog.WithInfoLevel(), golog.WithJSONEncoder())
	return logger, err
}

// Debug 日志
func Debug(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Debug(msg, fields...)
}

// Info 日志
func Info(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Info(msg, fields...)
}

// Warn 日志
func Warn(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Warn(msg, fields...)
}

// Error 日志
func Error(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Error(msg, fields...)
}

// Panic 日志
func Panic(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Panic(msg, fields...)
}

// Fatal 日志
func Fatal(tc *trace.Context, msg string, fields ...zap.Field) {
	if fields == nil {
		fields = make([]zap.Field, 0)
	}
	fields = append(fields, getTraceIDField(tc))
	golog.Fatal(msg, fields...)
}

// getTraceIDField get traceId
func getTraceIDField(tc *trace.Context) zap.Field {
	var field zap.Field
	if tc != nil {
		field = golog.Object(trace.TraceIDKey, tc.GetTraceID())
	} else {
		field = golog.String(trace.TraceIDKey, "")
	}

	return field
}
