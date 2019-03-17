// Package prometheus unit test
// Created by chenguolin 2019-03-16
package prometheus

import "testing"

func TestNewHistogram(t *testing.T) {
	name := "test_new_histogram"
	desc := "test new histogram metrics"
	buckets := []float64{1, 2, 3, 4, 5, 6.0, 7.0}
	labelKeys := []string{"name"}
	histogram := NewHistogram(name, desc, buckets, labelKeys)
	if histogram == nil {
		t.Fatal("TestNewHistogram histogram == nil")
	}
}

func TestHistogram_Observe(t *testing.T) {
	name := "test_new_histogram"
	desc := "test new histogram metrics"
	buckets := []float64{1, 2, 3, 4, 5, 6.0, 7.0}
	labelKeys := []string{"name"}
	histogram := NewHistogram(name, desc, buckets, labelKeys)

	labels := make(Labels)
	labels["name"] = "TestHistogram_Observe"
	histogram.Observe(labels, 1.0)
}
