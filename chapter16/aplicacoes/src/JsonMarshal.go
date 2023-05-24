package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	Conta     float64
}

func main() {
	jamesBond := pessoa{
		Nome:      "James",
		Sobrenome: "Bond",
		Idade:     40,
		Profissao: "espiao",
		Conta:     30.000,
	}

	dartVader := pessoa{
		Nome:      "Anakin",
		Sobrenome: "SkyWalker",
		Idade:     500,
		Profissao: "Jedi",
		Conta:     999999999.90,
	}

	j, err := json.Marshal(jamesBond) // transformando os strucs em json
	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(dartVader) // transformando os strucs em json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j)) //mostrando o json depois de passar de byte por string
	fmt.Println(string(d))
}
