package main

import (
	"fmt"
)

func main() {
	// --> If com condicional else
	//if x := 120; x > 100 {
	//	fmt.Println("x é maior que 100")
	//} else {
	//	fmt.Println("x é menor que 100")
	//}

	// --> If com condicional else if
	if x := 20; x > 100 {
		fmt.Println("x é maior que 100")
	} else if x < 10 {
		fmt.Println("x é menor que 10")
	} else {
		fmt.Println("x é menor que 100 e maior que 10")
	}
}
