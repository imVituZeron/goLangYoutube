package main

import (
	"fmt"
)

func main() {
	x := 10

	//if x < 11 { -> com verdade
	//if (x < 11) { -> outro modo de uma condicional if
	if !(x > 100) { // trocando o valor de false para true com o operado !
		fmt.Println(x)
	}

	if y := 20; y > 10 { // declarando o variavel e ja colocando o condiciona IF
		fmt.Println("y é maior que 10")
	}
}
