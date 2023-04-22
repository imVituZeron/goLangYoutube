package main

import (
	"fmt"
)

const x int = 10

const y = 10.0
const z = 4

func main() {
	fmt.Printf("valor: %v, tipo: %T\n", x, x)
	fmt.Printf("valor: %v, tipo: %T\n", y, y)
	fmt.Printf("valor: %v, tipo: %T\n", z, z)
}
