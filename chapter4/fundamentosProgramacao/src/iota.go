package main

import "fmt"

// iota é um tipo integer sequenciamente, onde que cada variavel que for atribuido o tipo
// vai receber um numero sequencial.
const (
	a = iota
	_ = iota
	c = iota
	_ = iota
	x = iota
	z = iota
)

func main() {
	fmt.Println(a, c, x, z)
}
