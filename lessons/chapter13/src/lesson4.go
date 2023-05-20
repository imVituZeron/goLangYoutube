package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) showInfo() {
	fmt.Printf("O nome completo é %v %v\n", p.nome, p.sobrenome)
	fmt.Printf("e a idade é %v\n", p.idade)
}

func main() {
	robert := pessoa{"Rober", "Cal", 34}
	robert.showInfo()
}
