package main

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	tracing, closer, err := config.Configuration{
		ServiceName: "hello.service",
		Sampler:     &config.SamplerConfig{Type: "const", Param: 1},
		Reporter:    &config.ReporterConfig{CollectorEndpoint: "http://localhost:14268/api/traces"},
	}.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracing)
	// web 示例
	http.HandleFunc("/api1", jaegerTracing(http.HandlerFunc(api1)))
	http.HandleFunc("/api2", jaegerTracing(http.HandlerFunc(api2)))
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}

// 添加jaeger 分布式追踪中间件
func jaegerTracing(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span := opentracing.StartSpan(r.URL.String())
		defer span.Finish()
		handler.ServeHTTP(w, r)
	}
}
