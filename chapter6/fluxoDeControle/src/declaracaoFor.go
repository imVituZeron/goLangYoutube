package main

import (
	"fmt"
)

func main() {
	// for x := 0; x < 10; x++ {
	// 	fmt.Println("asdfaf")
	// }  FOR tradicional

	x := 0

	for x < 10 {
		fmt.Println(x) // FOR imitando um while
		x++
	}

	for x < 10 {
		fmt.Println(x) // FOR infinito
	}

	for { // Um FOR infinito mas com uma condição de break
		if x < 10 {
			fmt.Println("X menos que 10")
		} else {
			fmt.Println("X maior que 10 - cabou")
			break
		}
	}

}
