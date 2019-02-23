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

// Exists check path exists
func Exists(path string) bool {
	// os.Stat get file info
	_, err := os.Stat(path)
	if err != nil {
		// check file exist error
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

// IsDir check path is directory
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

// IsFile check path is file
func IsFile(path string) bool {
	return !IsDir(path)
}
