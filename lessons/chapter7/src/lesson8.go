package main

import (
	"fmt"
)

func main() {
	pessoa := "Jorge marido da Anfisa"
	switch {
	case pessoa == "Sara":
		fmt.Println("Sara vai lavar a louça")
	case pessoa == "Vitor":
		fmt.Println("Vitor vai lavar a louça")
	default:
		fmt.Println("Qualquer um vai lavar a louça")
	}
}
