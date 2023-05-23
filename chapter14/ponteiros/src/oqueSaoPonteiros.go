package main

import (
	"fmt"
)

func main() {
	// o & mostra o endereco de uma variavel
	// o * mostra o conteudo do endereco de uma variavel

	x := 10

	y := &x // y esta armazenando o endereco de x

	fmt.Println(*y) // aqui estamos vendo o conteudo que tem no endereco de x
}
