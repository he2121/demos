package main

import (
	"fmt"
	"net/http"
)

func api1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello api1")
}

func api2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello api2")
}