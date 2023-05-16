package main

import "fmt"

func main() {
	x := 456

	y := func(x int) { // consegue passar a funcao para uma variavel
		fmt.Println(x, "vezes 45:")
		fmt.Println(x * 45)
	}

	y(x) // chamando a variavel comum argumento
}
