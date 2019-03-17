// Package prometheus unit test
// Created by chenguolin 2019-03-16
package prometheus

import (
	"testing"
)

func TestNewCounter(t *testing.T) {
	name := "test_new_counter"
	desc := "test new counter metrics"
	labelKeys := []string{"name"}
	counter := NewCounter(name, desc, labelKeys)
	if counter == nil {
		t.Fatal("TestNewCounter counter == nil")
	}
}

func TestCounter_Add(t *testing.T) {
	name := "test_new_counter"
	desc := "test new counter metrics"
	labelKeys := []string{"name"}
	counter := NewCounter(name, desc, labelKeys)

	labels := make(Labels)
	labels["name"] = "TestCounter_Add"
	counter.Add(labels, 1.0)
}
