// Package prometheus golang prometheus
// Created by chenguolin 2019-03-16
package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Counter metrics
type Counter struct {
	prometheus.CounterVec
}

// NewCounter new Counter metrics
func NewCounter(name, desc string, labels []string) *Counter {
	// new CounterVec
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: desc,
	}, labels)

	// register
	prometheus.MustRegister(counterVec)

	return &Counter{
		*counterVec,
	}
}

// Add counter metrics value
func (c *Counter) Add(labels prometheus.Labels, val float64) {
	c.With(labels).Add(val)
}
