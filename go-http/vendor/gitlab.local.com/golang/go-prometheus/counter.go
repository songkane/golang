// Package prometheus counter metrics
// Created by chenguolin 2019-03-16
package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Counter metrics
type Counter struct {
	counterVec *prometheus.CounterVec
}

// NewCounter new Counter metrics
func NewCounter(name, desc string, labelKeys []string) *Counter {
	// new CounterVec
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: desc,
	}, labelKeys)

	// register
	prometheus.MustRegister(counterVec)

	return &Counter{
		counterVec: counterVec,
	}
}

// Add counter metrics value
func (c *Counter) Add(labels Labels, val float64) {
	c.counterVec.With(prometheus.Labels(labels)).Add(val)
}
