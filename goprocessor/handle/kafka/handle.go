// Package kafka handle process
// Created by chenguolin 2019-01-13
package kafka

import (
	"fmt"

	"gitlab.local.com/golang/goprocessor/processor"
)

// Handle kafka handle
type Handle struct {
	// TODO
}

// NewHandle new kafka handle
func NewHandle() *Handle {
	return &Handle{}
}

// Process kafka record process
func (h *Handle) Process(record processor.Record) {
	// TODO record can convert 2 any type
	// TODO type message struct {
	// TODO    name string
	// TODO    age  int
	// TODO }
	// TODO msg, ok := record.(message)
	fmt.Println(record)
}
