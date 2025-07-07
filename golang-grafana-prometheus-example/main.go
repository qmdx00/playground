package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	collector = NewPrometheusCollector("biz")
)

func setupPrometheus() {
	collector.RegisterCounter("biz_foo_counter", []string{"method", "path"})
	collector.RegisterGauge("biz_foo_gauge", []string{"foo", "bar"})
}

func main() {
	setupPrometheus()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		collector.CounterVecWithLabelValues("biz_foo_counter", r.Method, r.URL.Path).Inc()
		_, _ = w.Write([]byte("foo called"))
	})

	go tick()

	http.ListenAndServe(":8080", nil)
}

func tick() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	for range ticker.C {
		// Simulate some metric updates
		collector.GaugeVecWithLabelValues("biz_foo_gauge", "foo_value", "bar_value").Set(rand.Float64())
	}
}
