// Package prometheus gauge metrics
// Created by chenguolin 2019-03-16
package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Gauge metrics
type Gauge struct {
	gaugeVec *prometheus.GaugeVec
}

// NewGauge new Gauge metrics
func NewGauge(name, desc string, labelKeys []string) *Gauge {
	// new GaugeVec
	gaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: desc,
	}, labelKeys)

	// register
	prometheus.MustRegister(gaugeVec)

	return &Gauge{
		gaugeVec: gaugeVec,
	}
}

// Set gauge value
func (g *Gauge) Set(label Labels, val float64) {
	g.gaugeVec.With(prometheus.Labels(label)).Set(val)
}
