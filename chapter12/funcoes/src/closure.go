package main

import (
	"fmt"
)

func main() {
	a := i() // chamada 1
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i() // chamada 2
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func i() func() int {
	x := 0
	// essa funcao que e o retorno da outra é um closure
	// ela pega uma variavel que esta fora do seu escopo
	// assim ela zera toda vez que e chamanda por outra variavel
	return func() int {
		x++
		return x
	}
}
