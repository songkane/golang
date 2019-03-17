// Package prometheus unit test
// Created by chenguolin 2019-03-16
package prometheus

import "testing"

func TestNewGauge(t *testing.T) {
	name := "test_new_gauge"
	desc := "test new gauge metrics"
	labelKeys := []string{"name"}
	gauge := NewGauge(name, desc, labelKeys)
	if gauge == nil {
		t.Fatal("TestNewGauge gauge == nil")
	}
}

func TestGauge_Set(t *testing.T) {
	name := "test_new_gauge"
	desc := "test new gauge metrics"
	labelKeys := []string{"name"}
	gauge := NewGauge(name, desc, labelKeys)

	labels := make(Labels)
	labels["name"] = "TestGauge_Set"
	gauge.Set(labels, 1.0)
}
