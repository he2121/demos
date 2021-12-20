package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 初始化 prometheus exporter
	initPrometheus()

	// web 示例
	http.HandleFunc("/api1", prometheusMetric(http.HandlerFunc(api1)))
	http.HandleFunc("/api2", prometheusMetric(http.HandlerFunc(api2)))
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}

func initPrometheus() {
	// 注册自定义的监控指标
	prometheus.MustRegister(apiTotalCounter)
	prometheus.MustRegister(apiCounterVec)
	prometheus.MustRegister(apiHandleMS)
	// 暴露 exporter 地址，prometheus server 通过 pull 这个地址，拉取指标数据
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(":1235", nil)
	}()
}

// metric 中间件
func prometheusMetric(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 总接口访问量计数
		apiTotalCounter.Inc()
		// 单个接口访问量计数
		apiCounterVec.WithLabelValues(r.URL.String()).Inc()
		start := time.Now()
		handler.ServeHTTP(w, r)
		ms := time.Now().UnixMicro() - start.UnixMicro()
		// 接口时延计数
		apiHandleMS.WithLabelValues(r.URL.String()).Observe(float64(ms))
	}
}
