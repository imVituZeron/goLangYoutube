package main

import (
	"fmt"
)

var x = 10
var y = 20

func main() {

	igual := x == y
	diferente := x != y
	maiorIgual := x >= y
	maior := x > y
	menorIgual := x <= y
	menor := x < y

	fmt.Printf("%v == %v: %v\n", x, y, igual)
	fmt.Printf("%v != %v: %v\n", x, y, diferente)
	fmt.Printf("%v >= %v: %v\n", x, y, maiorIgual)
	fmt.Printf("%v > %v: %v\n", x, y, maior)
	fmt.Printf("%v <= %v: %v\n", x, y, menorIgual)
	fmt.Printf("%v < %v: %v\n", x, y, menor)
}
