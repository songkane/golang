// Package prometheus golang prometheus
// Created by chenguolin 2019-03-16
package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Summary metrics
type Summary struct {
	prometheus.SummaryVec
}

// NewSummary new summary metrics
func NewSummary(name, desc string, labels []string) *Summary {
	// new HistogramVec
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: name,
		Help: desc,
	}, labels)

	// register
	prometheus.MustRegister(summaryVec)

	return &Summary{
		*summaryVec,
	}
}

// Observe metrics
func (s *Summary) Observe(label prometheus.Labels, val float64) {
	s.With(label).Observe(val)
}
