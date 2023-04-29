package main

import (
	"fmt"
)

func main() {
	x := 2000
	fmt.Println("Ano que eu nasce: ", x)
	for x <= 2023 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Ano atual: ", x-1)
}
