package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("O seu OS é:\t", runtime.GOOS)
	fmt.Println("O seu arch é:\t", runtime.GOARCH)
}
