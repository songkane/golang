// Package prometheus golang prometheus
// Created by chenguolin 2019-03-16
package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Histogram metrics
type Histogram struct {
	prometheus.HistogramVec
}

// NewHistogram new Histogram metrics
func NewHistogram(name, desc string, buckets []float64, labels []string) *Histogram {
	// new HistogramVec
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    desc,
		Buckets: buckets,
	}, labels)

	// register
	prometheus.MustRegister(histogramVec)

	return &Histogram{
		*histogramVec,
	}
}

// Observe metrics value
func (h *Histogram) Observe(label prometheus.Labels, val float64) {
	h.With(label).Observe(val)
}
