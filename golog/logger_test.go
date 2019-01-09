// Package golog unit test
// Created by chenguolin 2018-12-26
package golog

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// case 1
	log, err := NewLogger()
	if err != nil {
		t.Fatal("TestNewLogger case 1 NewLogger != nil")
	}
	if log == nil {
		t.Fatal("TestNewLogger case 1 log == nil")
	}

	// case 2
	log, err = NewLogger(WithOutput(os.Stdout), WithInfoLevel())
	if err != nil {
		t.Fatal("TestNewLogger case 2 NewLogger != nil")
	}
	if log == nil {
		t.Fatal("TestNewLogger case 2 log == nil")
	}

	// case 3
	writer, _ := NewRotateWriter(accessLog, "20060102-15")
	log, err = NewLogger(WithOutput(writer), WithInfoLevel())
	if err != nil {
		t.Fatal("TestNewLogger case 3 NewLogger != nil")
	}
	if log == nil {
		t.Fatal("TestNewLogger case 3 log == nil")
	}
}

func TestLogger_Debug(t *testing.T) {
	log, _ := NewLogger()
	log.Debug("TestLogger_Debug case", String("case", "TestLogger_Debug"))
}

func TestLogger_Info(t *testing.T) {
	log, _ := NewLogger()
	log.Info("TestLogger_Info case", String("case", "TestLogger_Info"))
}

func TestLogger_Warn(t *testing.T) {
	log, _ := NewLogger()
	log.Warn("TestLogger_Warn case", String("case", "TestLogger_Warn"))
}

func TestLogger_Error(t *testing.T) {
	log, _ := NewLogger()
	log.Error("TestLogger_Error case", String("case", "TestLogger_Error"))
}

/*
func TestLogger_Panic(t *testing.T) {
	log, _ := NewLogger()
	log.Panic("TestLogger_Panic case", String("case", "TestLogger_Panic"))
}

func TestLogger_Fatal(t *testing.T) {
	log, _ := NewLogger()
	log.Fatal("TestLogger_Fatal case", String("case", "TestLogger_Fatal"))
}
*/
