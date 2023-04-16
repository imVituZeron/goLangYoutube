package main

import "fmt"

type star int

var x star

func main() {
	fmt.Printf("%v %T\n", x, x) // valor e tipo de X
	x = 42
	fmt.Println(x)
}
