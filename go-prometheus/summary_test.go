// Package prometheus unit test
// Created by chenguolin 2019-03-16
package prometheus

import "testing"

func TestNewSummary(t *testing.T) {
	name := "test_new_summary"
	desc := "test new summary metrics"
	labelKeys := []string{"name"}
	histogram := NewSummary(name, desc, labelKeys)
	if histogram == nil {
		t.Fatal("TestNewSummary histogram == nil")
	}
}

func TestSummary_Observe(t *testing.T) {
	name := "test_new_summary"
	desc := "test new summary metrics"
	labelKeys := []string{"name"}
	summary := NewSummary(name, desc, labelKeys)

	labels := make(Labels)
	labels["name"] = "TestSummary_Observe"
	summary.Observe(labels, 1.0)
}
