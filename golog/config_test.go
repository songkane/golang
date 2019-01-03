/*
Package golog unit test
Created by chenguolin 2018-12-25
*/

package golog

import (
	"os"
	"testing"
)

func TestWithOutput(t *testing.T) {
	// case 1
	op := WithOutput(os.Stdout)
	if op == nil {
		t.Fatal("TestWithOutput case 1 op == nil")
	}

	// case 2
	writer, _ := NewRotateWriter(accessLog, "20060102-15")
	op = WithOutput(writer)
	if op == nil {
		t.Fatal("TestWithOutput case 2 op == nil")
	}
}

func TestWithNoLock(t *testing.T) {
	op := WithNoLock()
	if op == nil {
		t.Fatal("TestWithNoLock case 1 op == nil")
	}
}

func TestWithCaller(t *testing.T) {
	op := WithCaller()
	if op == nil {
		t.Fatal("TestWithCaller case 1 op == nil")
	}
}

func TestWithJSONEncoder(t *testing.T) {
	op := WithJSONEncoder()
	if op == nil {
		t.Fatal("TestWithJSONEncoder case 1 op == nil")
	}
}

func TestWithConsoleEncoder(t *testing.T) {
	op := WithConsoleEncoder()
	if op == nil {
		t.Fatal("TestWithConsoleEncoder case 1 op == nil")
	}
}

func TestWithDebugLevel(t *testing.T) {
	op := WithDebugLevel()
	if op == nil {
		t.Fatal("WithDebugLevel case 1 op == nil")
	}
}

func TestWithInfoLevel(t *testing.T) {
	op := WithInfoLevel()
	if op == nil {
		t.Fatal("WithInfoLevel case 1 op == nil")
	}
}

func TestWithWarnLevel(t *testing.T) {
	op := WithWarnLevel()
	if op == nil {
		t.Fatal("WithWarnLevel case 1 op == nil")
	}
}

func TestWithErrorLevel(t *testing.T) {
	op := WithErrorLevel()
	if op == nil {
		t.Fatal("WithErrorLevel case 1 op == nil")
	}
}

func TestWithPanicLevel(t *testing.T) {
	op := WithPanicLevel()
	if op == nil {
		t.Fatal("WithPanicLevel case 1 op == nil")
	}
}

func TestWithFatalLevel(t *testing.T) {
	op := WithFatalLevel()
	if op == nil {
		t.Fatal("WithFatalLevel case 1 op == nil")
	}
}

func TestWithTimePattern(t *testing.T) {
	op := WithTimePattern(DefaultTimePattern)
	if op == nil {
		t.Fatal("WithTimePattern case 1 op == nil")
	}
}
