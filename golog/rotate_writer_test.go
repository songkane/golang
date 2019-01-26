// Package golog unit test
// Created by chenguolin 2018-12-26
package golog

import (
	"testing"
	"time"
)

func TestNewRotateWriter(t *testing.T) {
	// case 1
	fileName := ""
	pattern := ""
	writer, err := NewRotateWriter(fileName, pattern)
	if err == nil {
		t.Fatal("TestNewRotateWriter case 1 err == nil")
	}
	if writer != nil {
		t.Fatal("TestNewRotateWriter case 1 writer != nil")
	}

	// case 2
	fileName = ""
	pattern = DefaultTimePattern
	writer, err = NewRotateWriter(fileName, pattern)
	if err == nil {
		t.Fatal("TestNewRotateWriter case 2 err == nil")
	}
	if writer != nil {
		t.Fatal("TestNewRotateWriter case 2 writer != nil")
	}

	// case 3
	fileName = accessLog
	pattern = ""
	writer, err = NewRotateWriter(fileName, pattern)
	if err == nil {
		t.Fatal("TestNewRotateWriter case 3 err == nil")
	}
	if writer != nil {
		t.Fatal("TestNewRotateWriter case 3 writer != nil")
	}

	// case 4
	fileName = "./test/logs/unit_test_log"
	pattern = DefaultTimePattern
	writer, err = NewRotateWriter(fileName, pattern)
	if err != nil {
		t.Fatal("TestNewRotateWriter case 4 err != nil")
	}
	if writer == nil {
		t.Fatal("TestNewRotateWriter case 4 writer == nil")
	}
}

func TestRotateWriter_Write(t *testing.T) {
	// case 1
	fileName := "./unit_test_log"
	pattern := DefaultTimePattern
	writer, _ := NewRotateWriter(fileName, pattern)

	// case 1
	size, err := writer.Write([]byte(""))
	if err != nil {
		t.Fatal("TestRotateWriter_Write case 1 writer.Write err != nil", err)
	}
	if size != 0 {
		t.Fatal("TestRotateWriter_Write case 1 writer.Write size != 0")
	}

	// case 2
	size, err = writer.Write([]byte("TestRotateWriter_Write case"))
	if err != nil {
		t.Fatal("TestRotateWriter_Write case 2 writer.Write err != nil", err)
	}
	if size <= 0 {
		t.Fatal("TestRotateWriter_Write case 2 writer.Write size <= 0")
	}
}

func TestRotateWriter_checkRoll(t *testing.T) {
	// case 1
	fileName := "./unit_test_log"
	pattern := DefaultTimePattern
	writer, _ := NewRotateWriter(fileName, pattern)

	preCheckPoint := writer.checkPoint
	preFp := writer.fp

	// time sleep 1 second
	time.Sleep(1 * time.Second)

	err := writer.checkRoll()
	if err != nil {
		t.Fatal("TestRotateWriter_checkRoll writer.checkRoll error", err)
	}
	curCheckPoint := writer.checkPoint
	curFp := writer.fp

	if preCheckPoint == curCheckPoint {
		t.Fatal("TestRotateWriter_checkRoll preCheckPoint == curCheckPoint")
	}
	if preFp == curFp {
		t.Fatal("TestRotateWriter_checkRoll preFp == curFp")
	}
}
