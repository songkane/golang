// Package sql hash unit test
// Created by chenguolin 2019-02-23
package sql

import (
	"fmt"
	"testing"
)

func TestCrc32(t *testing.T) {
	// case 1
	v := Crc32(0)
	if v != int64(4108050209) {
		t.Fatal("TestCrc32 case 1 failed ~")
	}

	// case 2
	v = Crc32(int64(1 << 31))
	if v != int64(2238651201) {
		t.Fatal("TestCrc32 case 2 failed ~")
	}

	// case 3
	v = Crc32(int64(1 << 60))
	fmt.Println(v)
	if v != int64(4271896104) {
		t.Fatal("TestCrc32 case 3 failed ~")
	}
}

func TestGetDbIndex(t *testing.T) {
	// case 1
	idx := GetDbIndex(int64(1<<31), 8)
	if idx != 1 {
		t.Fatal("TestGetDbIndex case 1 idx != 1")
	}

	// case 2
	idx = GetDbIndex(int64(1<<60), 8)
	if idx != 0 {
		t.Fatal("TestGetDbIndex case 2 idx != 0")
	}
}

func TestGetTableIndex(t *testing.T) {
	// case 1
	idx := GetTableIndex(int64(1<<31), 1024)
	if idx != 833 {
		t.Fatal("TestGetDbIndex case 1 idx != 833")
	}

	// case 2
	idx = GetTableIndex(int64(1<<60), 1024)
	if idx != 552 {
		t.Fatal("TestGetDbIndex case 2 idx != 552")
	}
}
