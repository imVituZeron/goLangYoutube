package main

import (
	"fmt"
)

func main() {
	x := 0
	recebeValor(x)     // <-- esta funcao pega uma copia do variavel e coloca na funcao
	recebePonteiro(&x) // <-- esta funcao pega o endereco de memoria da variavel e coloca na funcao
	fmt.Println(x)
}

func recebeValor(x int) {
	x++ // <-- atribui mais 1 na copia da variavel
	fmt.Println(x)
}

func recebePonteiro(x *int) {
	*x++ // <-- atribui mais 1 no endereco da variavel
	fmt.Println(x)
}
