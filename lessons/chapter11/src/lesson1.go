package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Sara",
		sobrenome: "Stefani",
		sabores:   []string{"ninho trufado", "pitaia", "flocos"},
	}

	pessoa2 := Pessoa{
		nome:      "Vitor",
		sobrenome: "Santos",
		sabores:   []string{"chocolate", "Chocolate", "CHOCOLATEEEE"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, value := range pessoa1.sabores {
		fmt.Println("\t", value)
	}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, value := range pessoa2.sabores {
		fmt.Println("\t", value)
	}
}
