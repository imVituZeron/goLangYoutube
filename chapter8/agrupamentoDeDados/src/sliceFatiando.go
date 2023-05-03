package main

import "fmt"

func main() {
	sabores := []string{"Marguerita", "Frango", "Calabresa", "4 queijos"}

	fatia := sabores[0:2]
	fatia2 := sabores[2:len(sabores)] // passando o tamanha do slice com o len
	// O primeiro é onde começa o corte do slice e o segundo paramentro é onde termina
	// mas não pegando ele.

	fmt.Println(fatia)
	fmt.Println(fatia2)

	for x := 0; x < len(sabores); x++ {
		fmt.Println(sabores[x])
	}

	fmt.Println(sabores)
	sabores = append(sabores[:2], sabores[3:]...) // criando um slice novo sem o valor do index 2
	fmt.Println(sabores)                          // mostrando novos valores so slice
}
