package main

import (
	"fmt"
)

const (
	x = 10
	y = 100
	z = 20 // constantes não tem tipos, só são definidos no uso
)

var c int
var d float64

func main() {
	d = x
	c = y

	fmt.Printf("tipo e valor de d: %T/%v\n", d, d)
	fmt.Printf("tipo e valor de c: %T/%v\n", c, c)
	fmt.Printf("tipo e valor de z: %T/%v\n", z, z)
}
