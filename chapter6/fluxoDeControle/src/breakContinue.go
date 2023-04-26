package main

import (
	"fmt"
)

// o break qebra o loop inteiro
// o continue quebra uma certa iteraçao do um loop
func main() {
	x := 0
	for x <= 20 {
		x++
		if x%2 != 0 { // ele quebra a iteração do loop no meio, sempre que o numero é impar!
			continue // por isso que na execução tem numeros impares
		}
		fmt.Println(x)
	}
}
