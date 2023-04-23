package main

import (
	"fmt"
)

const x = iota + 2024
const y = iota + 2024
const w = iota + 2024
const z = iota + 2024

func main() {
	fmt.Printf("O proximo ano é: %v\n", x)
	fmt.Printf("O proximo ano é: %v\n", y)
	fmt.Printf("O proximo ano é: %v\n", w)
	fmt.Printf("O proximo ano é: %v\n", z)
}
