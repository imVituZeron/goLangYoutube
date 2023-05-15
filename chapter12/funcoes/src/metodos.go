package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "bom dia!")
}

func main() {
	mauricio := pessoa{"Mauricio", 30}
	mauricio.oibomdia()

	fmt.Println("adf")
}
