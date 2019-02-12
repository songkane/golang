// Package executor 重试执行封装函数
// Created by chenguolin 2018-11-16
package executor

import (
	"errors"
	"math/rand"
	"testing"
)

func TestRunUntilSuccess(t *testing.T) {
	f := func() (interface{}, error) {
		x := rand.Intn(4)
		if x == 2 {
			return x, nil
		}

		return x, errors.New("error")
	}

	name := "TestRunUntilSuccess"
	x := RunUntilSuccess(name, f)
	if x != 2 {
		t.Fatal("TestRunUntilSuccess failed ~")
	}
}

func TestRunUntilSuccessNoRes(t *testing.T) {
	f := func() error {
		x := rand.Intn(4)
		if x == 2 {
			return nil
		}

		return errors.New("error")
	}

	name := "TestRunUntilSuccess"
	RunUntilSuccessNoRes(name, f)
}
