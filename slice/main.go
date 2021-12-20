package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abs, _ := filepath.Abs("../hello")

	fmt.Println(abs)
	fmt.Println(filepath.Base("../hello.txt"))
}
