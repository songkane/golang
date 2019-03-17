// Package prometheus histogram metrics
// Created by chenguolin 2019-03-16
package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Histogram metrics
type Histogram struct {
	histogramVec *prometheus.HistogramVec
}

// NewHistogram new Histogram metrics
func NewHistogram(name, desc string, buckets []float64, labelKeys []string) *Histogram {
	// new HistogramVec
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    desc,
		Buckets: buckets,
	}, labelKeys)

	// register
	prometheus.MustRegister(histogramVec)

	return &Histogram{
		histogramVec: histogramVec,
	}
}

// Observe metrics value
func (h *Histogram) Observe(label Labels, val float64) {
	h.histogramVec.With(prometheus.Labels(label)).Observe(val)
}
