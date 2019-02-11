// Package file unit test
// Created by chenguolin 2018-02-11
package file

import (
	"os"
	"testing"
)

func TestDumpFile(t *testing.T) {
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	content := string("TestDumpFile")

	// case 1
	err := Dump("", []byte(content))
	if err.Error() != "open : no such file or directory" {
		t.Fatal("TestDumpFile case 1 failed ~")
	}

	// case 2
	err = Dump(fileName, []byte(content))
	if err != nil {
		t.Fatal("TestDumpFile case 2 failed ~")
	}
}

func TestLoadFile(t *testing.T) {
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	content := string("TestDumpFile")

	// case 1
	newCon, err := Load("")
	if err.Error() != "open : no such file or directory" {
		t.Fatal("TestLoadFile case 1 failed ~")
	}

	// case 2
	_ = Dump(fileName, []byte(content))
	newCon, err = Load(fileName)
	if err != nil {
		t.Fatal("TestLoadFile case 2 failed ~")
	}
	if string(newCon) != content {
		t.Fatal("TestLoadFile case 2 failed ~")
	}
}

func TestRemove(t *testing.T) {
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	content := string("TestDumpFile")

	// case 1
	err := Remove("")
	if err.Error() != "remove : no such file or directory" {
		t.Fatal("TestRemove case 1 failed ~")
	}

	// case 2
	_ = Dump(fileName, []byte(content))
	err = Remove(fileName)
	if err != nil {
		t.Fatal("TestRemove case 2 failed ~")
	}
}
