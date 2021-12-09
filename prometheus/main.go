package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	initPrometheus()

	http.HandleFunc("/api1", prometheusMetric(http.HandlerFunc(api1)))
	http.HandleFunc("/api2", prometheusMetric(http.HandlerFunc(api2)))

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
func initPrometheus() {
	prometheus.MustRegister(apiTotalCounter)
	prometheus.MustRegister(apiCounterVec)
	prometheus.MustRegister(apiHandleMS)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(":1235", nil)
	}()
}

// metric 中间件
func prometheusMetric(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiTotalCounter.Inc()
		apiCounterVec.WithLabelValues(r.URL.String()).Inc()
		start := time.Now()
		handler.ServeHTTP(w, r)
		ms := time.Now().UnixMicro() - start.UnixMicro()
		apiHandleMS.WithLabelValues(r.URL.String()).Observe(float64(ms))
	}
}
