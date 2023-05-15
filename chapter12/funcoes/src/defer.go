package main

import "fmt"

// o defer faz o a linha de codigo que ele tiver executar por ultimo.
func main() {
	// quando temos varios defer, o primeiro que entrar na lista vai ser o ultimo a ser executado.
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}
