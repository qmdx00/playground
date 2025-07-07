package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusCollector struct {
	namespace string

	histogram map[string]*prometheus.HistogramVec
	counter   map[string]*prometheus.CounterVec
	gauge     map[string]*prometheus.GaugeVec
	summary   map[string]*prometheus.SummaryVec
}

func NewPrometheusCollector(namespace string) *PrometheusCollector {
	return &PrometheusCollector{
		namespace: namespace,
		histogram: make(map[string]*prometheus.HistogramVec),
		counter:   make(map[string]*prometheus.CounterVec),
		gauge:     make(map[string]*prometheus.GaugeVec),
		summary:   make(map[string]*prometheus.SummaryVec),
	}
}

// RegisterHistogram ...
func (p *PrometheusCollector) RegisterHistogram(name string, labels []string) *prometheus.HistogramVec {
	if vec, exists := p.histogram[name]; exists {
		return vec
	}

	vec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: p.namespace,
		Name:      name,
	}, labels)

	p.histogram[name] = vec
	prometheus.MustRegister(vec)

	return vec
}

// RegisterCounter ...
func (p *PrometheusCollector) RegisterCounter(name string, labels []string) *prometheus.CounterVec {
	if vec, exists := p.counter[name]; exists {
		return vec
	}

	vec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: p.namespace,
		Name:      name,
	}, labels)

	p.counter[name] = vec
	prometheus.MustRegister(vec)

	return vec
}

// RegisterGauge ...
func (p *PrometheusCollector) RegisterGauge(name string, labels []string) *prometheus.GaugeVec {
	if vec, exists := p.gauge[name]; exists {
		return vec
	}

	vec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: p.namespace,
		Name:      name,
	}, labels)

	p.gauge[name] = vec
	prometheus.MustRegister(vec)

	return vec
}

// RegisterSummary ...
func (p *PrometheusCollector) RegisterSummary(name string, labels []string) *prometheus.SummaryVec {
	if vec, exists := p.summary[name]; exists {
		return vec
	}

	vec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: p.namespace,
		Name:      name,
	}, labels)

	p.summary[name] = vec
	prometheus.MustRegister(vec)

	return vec
}

func (p *PrometheusCollector) HistogramVecWithLabelValues(name string, labels ...string) prometheus.Observer {
	if vec, exists := p.histogram[name]; exists {
		return vec.WithLabelValues(labels...)
	}
	return nil
}

func (p *PrometheusCollector) CounterVecWithLabelValues(name string, labels ...string) prometheus.Counter {
	if vec, exists := p.counter[name]; exists {
		return vec.WithLabelValues(labels...)
	}
	return nil
}

func (p *PrometheusCollector) GaugeVecWithLabelValues(name string, labels ...string) prometheus.Gauge {
	if vec, exists := p.gauge[name]; exists {
		return vec.WithLabelValues(labels...)
	}
	return nil
}

func (p *PrometheusCollector) SummaryVecWithLabelValues(name string, labels ...string) prometheus.Observer {
	if vec, exists := p.summary[name]; exists {
		return vec.WithLabelValues(labels...)
	}
	return nil
}
