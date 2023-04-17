package main

import "fmt"

var x bool

func main() {
	fmt.Printf("x = %v\n", x)
	x = true
	fmt.Printf("x = %v\n", x)
	x = 10 < 100
	y := 10 == 100
	z := 10 > 100
	fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
}
