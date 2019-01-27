/*
Package logger 单元测试
Created by chenguolin 2018-11-18
*/
package logger

import (
	"testing"

	"gitlab.local.com/golang/go-common/trace"
	"gitlab.local.com/golang/go-log"
	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	tc := trace.WithTraceID("test")
	msg := "TestDebug output"
	fields := make([]zap.Field, 0)
	fields = append(fields, golog.Object("level", "debug"))
	Debug(tc, msg, fields...)
}

func TestWarn(t *testing.T) {
	tc := trace.WithTraceID("test")
	msg := "TestWarn output"
	fields := make([]zap.Field, 0)
	fields = append(fields, golog.Object("level", "warn"))
	Warn(tc, msg, fields...)
}

func TestInfo(t *testing.T) {
	tc := trace.WithTraceID("test")
	msg := "TestInfo output"
	fields := make([]zap.Field, 0)
	fields = append(fields, golog.Object("level", "info"))
	Info(tc, msg, fields...)
}

/*
func TestPanic(t *testing.T) {
	tc := trace.WithTraceID("test")
	msg := "TestPanic output"
	fields := make([]*log.Field, 0)
	fields = append(fields, log.Object("level", "panic"))
	Panic(tc, msg, fields...)
}

func TestFatal(t *testing.T) {
	tc := trace.WithTraceID("test")
	msg := "TestFatal output"
	fields := make([]*log.Field, 0)
	fields = append(fields, log.Object("level", "fatal"))
	Fatal(tc, msg, fields...)
}
*/
