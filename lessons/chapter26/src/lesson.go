package main

//

import "fmt"

func main() {
	fmt.Printf("Seu cachorro tem %v anos\n", Idade(4))
}

// A função Idade recebe um int e depois ele retona um int
func Idade(ano int) int {
	idade := ano * 7
	return idade
}
