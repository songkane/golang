// Package mysql mysql handle process
// Created by chenguolin 2019-01-13
package mysql

import (
	"fmt"

	"gitlab.local.com/golang/go-processor/processor"
)

// Handle mysql handle
type Handle struct {
	// TODO
}

// NewHandle new mysql handle
func NewHandle() *Handle {
	return &Handle{}
}

// Process mysql record process
func (h *Handle) Process(record processor.Record) {
	// TODO record can convert 2 any type
	// TODO type user struct {
	// TODO    name string
	// TODO    age  int
	// TODO }
	// TODO us, ok := record.(user)
	fmt.Println(record)

	// TODO 业务逻辑
}
