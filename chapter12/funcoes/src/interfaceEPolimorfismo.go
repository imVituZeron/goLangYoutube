package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type destista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type desenvolvedor struct {
	pessoa
	linguagempref string
	salario       float64
}

func (x destista) oibomdia() {
	fmt.Println("meu nome é", x.nome, "e eu ja arranquei", x.dentesarrancados, "olha, Bom dia!")
}

func (x desenvolvedor) oibomdia() {
	fmt.Println("meu nome é", x.nome, "e minha linguagem preferida é", x.linguagempref, "olha, Bom dia!")
}

type gente interface {
	oibomdia()
}

func serHumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case destista:
		fmt.Println("EU ganho:", g.(destista).salario)
	case desenvolvedor:
		fmt.Println("EU ganho:", g.(desenvolvedor).salario)
	}

}

func main() {
	dentista := destista{pessoa{"Jorge", "filipe", 34}, 12333, 40000.0}
	desenvolvedor := desenvolvedor{pessoa{"Vitor", "Santos", 24}, "goLang", 70000.0}

	serHumano(dentista)
	serHumano(desenvolvedor)
}
