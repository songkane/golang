// Package trace 单元测试
// Created by chenguolin 2018-11-18
package trace

import (
	"testing"
)

func TestDefaultTraceContext(t *testing.T) {
	tc := DefaultTraceContext()
	if tc == nil {
		t.Fatal("TestDefaultTraceContext DefaultTraceContext is nil")
	}
}

func TestWithTraceID(t *testing.T) {
	traceID := "test"
	tc := WithTraceID(traceID)
	if tc == nil {
		t.Fatal("TestWithTraceId WithTraceID tc == nil")
	}
}

func TestTraceContext_GetTraceID(t *testing.T) {
	// case 1
	tc := DefaultTraceContext()
	if tc == nil {
		t.Fatal("TestDefaultTraceContext DefaultTraceContext is nil")
	}
	traceID := tc.GetTraceID()
	if traceID != "" {
		t.Fatal("TestTraceContext_GetTraceId case 1 traceId != \"\"")
	}

	// case 2
	tc = WithTraceID("test")
	if tc == nil {
		t.Fatal("TestDefaultTraceContext WithTraceID is nil")
	}
	traceID = tc.GetTraceID()
	if traceID != "test" {
		t.Fatal("TestTraceContext_GetTraceId case 2 traceId != \"test\"")
	}
}

func TestTraceContext_SetTraceID(t *testing.T) {
	tc := DefaultTraceContext()
	if tc == nil {
		t.Fatal("TestDefaultTraceContext DefaultTraceContext is nil")
	}
	tc.SetTraceID("test")
	traceID := tc.GetTraceID()
	if traceID != "test" {
		t.Fatal("TestTraceContext_SetTraceId tc.GetTraceID != \"test\"")
	}
}
