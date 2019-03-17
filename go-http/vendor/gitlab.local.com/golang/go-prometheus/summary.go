// Package prometheus summary metrics
// Created by chenguolin 2019-03-16
package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Summary metrics
type Summary struct {
	summaryVec *prometheus.SummaryVec
}

// NewSummary new summary metrics
func NewSummary(name, desc string, labelKeys []string) *Summary {
	// new HistogramVec
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: name,
		Help: desc,
	}, labelKeys)

	// register
	prometheus.MustRegister(summaryVec)

	return &Summary{
		summaryVec: summaryVec,
	}
}

// Observe metrics
func (s *Summary) Observe(label Labels, val float64) {
	s.summaryVec.With(prometheus.Labels(label)).Observe(val)
}
