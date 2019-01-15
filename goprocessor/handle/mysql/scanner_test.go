// Package mysql scanner
// Created by chenguolin 2019-01-13
package mysql

import (
	"fmt"
	"testing"
	"time"

	"gitlab.local.com/golang/gomysql"
)

func TestNewScanner(t *testing.T) {
	scanner := NewScanner(1, time.Duration(5)*time.Second, &mysql.Mysql{})
	if scanner == nil {
		t.Fatal("TestNewScanner scanner == nil")
	}
	if scanner.isRunning != false {
		t.Fatal("TestNewScanner scanner.isRunning != false")
	}
	if scanner.maxChanSize != 1 {
		t.Fatal("TestNewScanner scanner.maxChanSize != 1")
	}
	if scanner.scanInterval != time.Duration(5)*time.Second {
		t.Fatal("TestNewScanner scanner.scanInterval != time.Duration(5)*time.Second")
	}
	if len(scanner.stopChan) != 0 {
		t.Fatal("TestNewScanner len(scanner.stopChan) != 0")
	}
	if len(scanner.records) != 0 {
		t.Fatal("TestNewScanner len(scanner.records) != 0")
	}
	if scanner.dbProxy == nil {
		t.Fatal("TestNewScanner scanner.dbProxy == nil")
	}
}

func TestScanner_Start(t *testing.T) {
	scanner := NewScanner(1, 5*time.Second, &mysql.Mysql{})
	scanner.Start()
	if scanner.isRunning != true {
		t.Fatal("TestScanner_Start scanner.isRunning != true")
	}
	time.Sleep(6 * time.Second)
	if len(scanner.records) == 0 {
		t.Fatal("TestScanner_Start len(scanner.records) == 0")
	}
}

func TestScanner_Stop(t *testing.T) {
	scanner := NewScanner(1, 5*time.Second, &mysql.Mysql{})
	scanner.Start()
	if scanner.isRunning != true {
		t.Fatal("TestScanner_Start scanner.isRunning != true")
	}
	time.Sleep(6 * time.Second)
	if len(scanner.records) == 0 {
		t.Fatal("TestScanner_Start len(scanner.records) == 0")
	}
	scanner.Stop()
	if scanner.isRunning != false {
		t.Fatal("TestScanner_Start scanner.isRunning != false")
	}
}

func TestScanner_Next(t *testing.T) {
	scanner := NewScanner(1, 5*time.Second, &mysql.Mysql{})
	scanner.Start()
	if scanner.isRunning != true {
		t.Fatal("TestScanner_Start scanner.isRunning != true")
	}
	time.Sleep(6 * time.Second)
	if len(scanner.records) == 0 {
		t.Fatal("TestScanner_Start len(scanner.records) == 0")
	}

	record, ok := scanner.Next()
	if !ok {
		t.Fatal("TestScanner_Start scanner.Next not ok")
	}
	fmt.Println(record)
}

func TestScanner_IsStopped(t *testing.T) {
	scanner := NewScanner(1, 5*time.Second, &mysql.Mysql{})
	scanner.Start()
	if scanner.IsStopped() == true {
		t.Fatal("TestScanner_IsStopped scanner.IsStopped() == true")
	}
	time.Sleep(1 * time.Second)
	scanner.Stop()
	if scanner.IsStopped() == false {
		t.Fatal("TestScanner_IsStopped scanner.IsStopped() == false")
	}
}
