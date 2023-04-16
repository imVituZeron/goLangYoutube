package main

import "fmt"

type star int

var x star
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	// conversao
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
