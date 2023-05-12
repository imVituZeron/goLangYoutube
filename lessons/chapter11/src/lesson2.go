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
	mepe := make(map[string]Pessoa)

	mepe2 := map[string]Pessoa{
		"Seu": Pessoa{"Jorge", "Seu", []string{"flocos", "cafe", "CAFEE"}},
	}

	mepe["Santos"] = Pessoa{"Vitor", "Santos", []string{"Sabao em po", "pe de porco"}}

	for key, value := range mepe {
		fmt.Println(key)
		fmt.Println(value.nome)
		for _, v := range value.sabores {
			fmt.Println("\t", v)
		}
	}

	for key, value := range mepe2 {
		fmt.Println(key)
		fmt.Println(value.nome)
		for _, v := range value.sabores {
			fmt.Println("\t", v)
		}
	}
}
