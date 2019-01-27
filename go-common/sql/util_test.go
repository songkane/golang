// Package sql mysql 基础功能函数
// Created by chenguolin 2018-05-21
package sql

import (
	"testing"
	"errors"
)

func TestIsDuplicateKeyError(t *testing.T) {
	// case 1
	e := errors.New("")
	ret := IsDuplicateKeyError(e)
	if ret != false {
		t.Fatal("TestIsDuplicateKeyError case 1 ret != false")
	}

	// case 2
	e = errors.New("error")
	ret = IsDuplicateKeyError(e)
	if ret != false {
		t.Fatal("TestIsDuplicateKeyError case 2 ret != false")
	}

	// case 3
	e = errors.New("Duplicate entry error")
	ret = IsDuplicateKeyError(e)
	if ret == false {
		t.Fatal("TestIsDuplicateKeyError case 3 ret != false")
	}
}

func TestIsNotRowsError(t *testing.T) {
	// case 1
	e := errors.New("")
	ret := IsNotRowsError(e)
	if ret != false {
		t.Fatal("TestIsNotRowsError case 1 ret != false")
	}

	// case 2
	e = errors.New("error")
	ret = IsNotRowsError(e)
	if ret != false {
		t.Fatal("TestIsNotRowsError case 2 ret != false")
	}

	// case 3
	e = errors.New("sql: no rows error")
	ret = IsNotRowsError(e)
	if ret == false {
		t.Fatal("TestIsNotRowsError case 3 ret != false")
	}
}
