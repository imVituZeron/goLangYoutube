package main

import (
	"fmt"
)

func main() {
	x := 2000
	for {
		fmt.Println(x)
		if x == 2023 {
			break
		}
		x++
	}
}
