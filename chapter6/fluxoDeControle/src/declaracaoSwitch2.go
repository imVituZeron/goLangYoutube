package main

import (
	"fmt"
)

var x interface{}

func main() {
	x = "aqui jas"

	switch x.(type) { // posso mandar fazer algo só pelo tipo da variavel
	case int:
		fmt.Println("e um Int")
	case bool:
		fmt.Println("e um bool")
	case string:
		fmt.Println("e uma string")
	case float64:
		fmt.Println("e um float")
	}

	switch j := 3; { // igual no for e no IF podemos iniciar a variavel no Switch também
	case j == 1:
		fmt.Println("e um 1")
	case j == 2:
		fmt.Println("e um 2")
	case j == 3:
		fmt.Println("e um 3")
	case j == 4:
		fmt.Println("e um 4")
	}
}
