package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(loop(6))
}

// usando a funcao pra chamar ela mesma ate chegar a 1
func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

// usando for
func loop(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
