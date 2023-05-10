package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct { // profissional recebe a pessoa (heranca)
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "alfredo",
		idade: 30,
	}

	pessoa2 := profissional{ // a melhor forma de instanciar uma struct dentro de outra
		pessoa: pessoa{
			nome:  "maria",
			idade: 31,
		},
		titulo:  "DevOps",
		salario: 1800,
	}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
