package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	sara := pessoa{"sara", "Stefani", 24}
	fmt.Println(sara)
	mudeMe(&sara)
	fmt.Println(sara)
}

func mudeMe(P *pessoa) {
	(*P).nome = "Vitor"
}
