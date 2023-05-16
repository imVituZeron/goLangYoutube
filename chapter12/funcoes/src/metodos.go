package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() { //comecamos com POO, esse metodo e do tipo pessoa
	fmt.Println(p.nome, "bom dia!")
}

func main() {
	mauricio := pessoa{"Mauricio", 30}
	mauricio.oibomdia()

	//oibomdia() //nao funcionar pois nao tem nenhuma instacia do tipo pessoa para usar esse metodo.

	fmt.Println("adf")
}
