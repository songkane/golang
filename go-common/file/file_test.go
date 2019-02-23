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

func TestExists(t *testing.T) {
	// case 1
	suc := Exists("")
	if suc == true {
		t.Fatal("TestExists case 1 suc == true")
	}

	// case 2
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	suc = Exists(fileName)
	if suc == true {
		t.Fatal("TestExists case 2 suc == true")
	}

	// case 3
	fileName = dir + "/file.go"
	suc = Exists(fileName)
	if suc != true {
		t.Fatal("TestExists case 3 suc != true")
	}
}

func TestIsDir(t *testing.T) {
	// case 1
	suc := IsDir("")
	if suc == true {
		t.Fatal("TestIsDir case 1 suc == true")
	}

	// case 2
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	suc = IsDir(fileName)
	if suc == true {
		t.Fatal("TestIsDir case 2 suc == true")
	}

	// case 3
	suc = IsDir(dir)
	if suc != true {
		t.Fatal("TestIsDir case 3 suc != true")
	}
}

func TestIsFile(t *testing.T) {
	// case 1
	suc := IsFile("")
	if suc != true {
		t.Fatal("IsFile case 1 suc != true")
	}

	// case 2
	dir, _ := os.Getwd()
	fileName := dir + "/test.dat"
	suc = IsFile(fileName)
	if suc != true {
		t.Fatal("IsFile case 2 suc != true")
	}

	// case 3
	suc = IsFile(dir)
	if suc == true {
		t.Fatal("IsFile case 3 suc == true")
	}
}
