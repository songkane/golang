/*
Package golog unit test
Created by chenguolin 2018-12-26
*/
package golog

import (
	"os"
	"testing"
)

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	if log == nil {
		t.Fatal("TestGetLogger GetLogger is nil")
	}
	if log.Log == nil {
		t.Fatal("TestGetLogger log.Log == nil")
	}
}

func TestSetLogger(t *testing.T) {
	// case 1
	SetLogger(nil)
	if GetLogger() != nil {
		t.Fatal("TestSetLogger case 1 GetLogger() != nil")
	}

	// case 2
	log, err := NewLogger(WithOutput(os.Stdout))
	if err != nil {
		t.Fatal("TestSetLogger case 2 NewLogger error", err)
	}
	SetLogger(log)
	if GetLogger() == nil {
		t.Fatal("TestSetLogger case 2 GetLogger() == nil")
	}
}

func TestDebug(t *testing.T) {
	Debug("TestDebug case")
}

func TestInfo(t *testing.T) {
	Info("TestInfo case")
}

func TestWarn(t *testing.T) {
	Warn("TestWarn case")
}

func TestError(t *testing.T) {
	Error("TestError case")
}

/*
func TestPanic(t *testing.T) {
	Panic("TestPanic case")
}

func TestFatal(t *testing.T) {
	Fatal("TestFatal case")
}
*/
