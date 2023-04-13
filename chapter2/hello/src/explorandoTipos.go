package main

import "fmt"

var x int
var y float32
var w float64
var z bool
var a string

// se não tiver atribuição de valor nas variaveis com escopo global
// só poderá haver essa atribuição dentro de algum escopo!
func main() {

	x = 10
	y = 11.1
	w = 111.11
	z = false
	a = "oi"

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("w:", w)
	fmt.Println("z:", z)
	fmt.Println("a:", a)
}
