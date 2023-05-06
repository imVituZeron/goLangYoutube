package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"Charles": 12341234,
		"Joana":   9996674,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["Joana"])

	amigos["Sara"] = 90123455 // modo de criar uma valor no map
	fmt.Println(amigos)
	fmt.Println(amigos["Sara"])

	fmt.Println(amigos["Vitoria"]) // mostrando valores que nao existem e recebendo um valor comma ok

	sera, ok := amigos["Jorge"] // mostrando valores que nao existem e recebendo um valor comma ok
	// comma ok idiom, ele meio que verifica se existe tao coisa, se existir o ok=true se nao ok=false

	fmt.Println(sera, ok)

}
