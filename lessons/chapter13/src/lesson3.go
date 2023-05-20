package main

import (
	"fmt"
)

// varios defer funcional igual uma pilha
func main() {
	defer fmt.Println("Quarto")
	defer fmt.Println("Terceiro")
	defer fmt.Println("Segundo")
	fmt.Println("Primeiro")
}
