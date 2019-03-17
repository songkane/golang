// Package prometheus golang prometheus
// Created by chenguolin 2019-03-16
package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Gauge metrics
type Gauge struct {
	prometheus.GaugeVec
}

// NewGauge new Gauge metrics
func NewGauge(name, desc string, labels []string) *Gauge {
	// new GaugeVec
	gaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: desc,
	}, labels)

	// register
	prometheus.MustRegister(gaugeVec)

	return &Gauge{
		*gaugeVec,
	}
}

// Set gauge value
func (g *Gauge) Set(label prometheus.Labels, val float64) {
	g.With(label).Set(val)
}
