// Package file 文件操作
// Created by chenguolin 2018-02-11
package file

import (
	"io/ioutil"
	"os"
)

// Dump dump content 2 file
func Dump(fileName string, content []byte) error {
	return ioutil.WriteFile(fileName, content, 0644)
}

// Load load file
func Load(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// Remove remove file
func Remove(fileName string) error {
	return os.Remove(fileName)
}
