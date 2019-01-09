// Package golog rotate writer
// Created by chenguolin 2018-12-26
package golog

import (
	"errors"
	"os"
	"sync"
	"time"
)

// RotateWriter writer
type RotateWriter struct {
	lock        sync.RWMutex //读写锁
	fileName    string       //文件名
	timePattern string       //时间格式
	checkPoint  string       //checkPoint用于校验是否需要切割文件
	fp          *os.File     //文件句柄
}

// NewRotateWriter return new RotateWriter
// @fileName target file name
// @pattern time pattern
func NewRotateWriter(fileName, pattern string) (*RotateWriter, error) {
	if fileName == "" || pattern == "" {
		return nil, errors.New("fileName or pattern invalid")
	}

	w := &RotateWriter{
		fileName:    fileName,
		timePattern: pattern,
		checkPoint:  time.Now().Format(pattern),
	}

	// Create creates the named file with mode 0666 (before umask), truncating
	// it if it already exists. If successful, methods on the returned
	// File can be used for I/O; the associated file descriptor has mode O_RDWR.
	// If there is an error, it will be of type *PathError.
	fp, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	w.fp = fp
	return w, nil
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
	// check need rotate
	err := w.checkRoll()
	if err != nil {
		return 0, err
	}

	// write 2 file
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.fp.Write(output)
}

// checkRoll check need rotate file
func (w *RotateWriter) checkRoll() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	nowPoint := time.Now().Format(w.timePattern)
	// No Need for Rotate
	if nowPoint == w.checkPoint {
		return nil
	}

	// Close existing file if open
	if w.fp != nil {
		err := w.fp.Close()
		w.fp = nil
		if err != nil {
			return err
		}
	}

	// Rename target file
	_, err := os.Stat(w.fileName)
	if err == nil {
		err = os.Rename(w.fileName, w.fileName+"."+w.checkPoint)
		if err != nil {
			return err
		}
	}

	// Create a file.
	w.fp, err = os.Create(w.fileName)
	w.checkPoint = nowPoint

	return nil
}
