package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
)

func api1(w http.ResponseWriter, r *http.Request) {
	// 根 ctx
	ctx := context.Background()
	// 记录 sever1.api 的调动耗时等信息
	span, c := opentracing.StartSpanFromContext(ctx, "server1.api1")
	defer span.Finish()
	fmt.Println("hello api1")
	time.Sleep(time.Second)
	// 模拟一个跨服务的调用，ctx 上下文信息传递
	service2XXX(c)
}

func service2XXX(ctx context.Context) {
	// 从 ctx 获取 span， 这样就能把在一条调用链的 span 聚合在一起
	span, c := opentracing.StartSpanFromContext(ctx, "server2.XXX")
	defer span.Finish()
	time.Sleep(time.Second * 2)
	fmt.Println("hello server2 xxx", c)
}

func api2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello api2")
}
