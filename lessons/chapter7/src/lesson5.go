package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("número: %v, resultado: %v\n", x, x%4)
	}
}
