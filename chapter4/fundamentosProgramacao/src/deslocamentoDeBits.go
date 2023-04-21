package main

import "fmt"

func main() {

	x := 1
	y := x << 1 // ele vai deslocar a quantidade de bits que voce pedir
	z := y << 1
	a := z << 5
	b := a >> 2 // aqui o operador esta voltando bits.

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
}
