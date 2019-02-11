// Package trace trace定义 用于上下文关联，常见于log打印traceId使用
// Created by chenguolin 2018-11-18
package trace

import (
	"time"
)

// Context 上下文定义结构体
type Context struct {
	values map[interface{}]interface{}
}

const (
	// TraceIDKey 常量定义
	TraceIDKey = "trace_id"
)

// DefaultTraceContext 获取默认Context实例
func DefaultTraceContext() *Context {
	return &Context{
		values: make(map[interface{}]interface{}),
	}
}

// WithTraceID 根据trace ID获取Context实例
func WithTraceID(traceID string) *Context {
	values := make(map[interface{}]interface{})
	values[TraceIDKey] = traceID
	return &Context{
		values: values,
	}
}

// GetTraceID 获取trace ID
func (tc *Context) GetTraceID() string {
	if value, ok := tc.values[TraceIDKey]; ok {
		return value.(string)
	}
	return string("")
}

// SetTraceID 设置trace ID
func (tc *Context) SetTraceID(traceID string) {
	tc.values[TraceIDKey] = traceID
}

// Deadline 默认实现
func (tc *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done 默认实现
func (tc *Context) Done() <-chan struct{} {
	return nil
}

// Err 默认实现
func (tc *Context) Err() error {
	return nil
}

// Value 根据key获取value
func (tc *Context) Value(key interface{}) interface{} {
	return tc.values[key]
}
