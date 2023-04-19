package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e" // todos do tipo string mais com valores em bytes diferentes como mostra nas linhas a baixo
	b := "é"
	c := "\u2312"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v\n", d, e, f)

	fmt.Println("--------------------------------")

	x := 10
	y := 10.0
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Println("--------------------------------")

	fmt.Println(runtime.GOOS)   // traz o SO
	fmt.Println(runtime.GOARCH) // Traz a arquitetura do computador.
}
