package main

import "github.com/prometheus/client_golang/prometheus"

var apiCounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_api_counter",
	Help: "web http requests number",
}, []string{"handler"})

var apiTotalCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_api_total_counter",
	Help: "web http requests number",
})

var apiHandleMS = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_api_handler_ms",
	Help:    "Microseconds of HTTP interface processing",
	Buckets: prometheus.LinearBuckets(1, 19, 10),
}, []string{"handler"})
