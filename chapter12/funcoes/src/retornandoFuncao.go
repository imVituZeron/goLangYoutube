package main

import "fmt"

func main() {
	// x vai receber o retorno do retornaFuncao e que é a outra funcao
	x := retornaFuncao()
	// chamando a outra funcao e esperando o retorno que é i * 10
	y := x(2)
	fmt.Println(y)

	//chamando a funcao e ja passando argumentos para o seu retorno que também é uma funcao
	fmt.Println(retornaFuncao()(4))
}

// func NOME(paramentros) RETORNO(parementros para o retorno) RETORNO do RETORNO {}
func retornaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
