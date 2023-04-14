package main

import "fmt"

type hotdog int // criando um tipo novo

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b) // as duas tem tipo diferentes

	fmt.Printf("%v\n", x) // as duas tem os mesmo valor
	fmt.Printf("%v\n", b)

	// x = b // não são as mesmas coisas

	fmt.Println("type is life")
}
