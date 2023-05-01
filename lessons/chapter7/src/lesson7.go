package main

import "fmt"

func main() {
	x := 1000
	if x == 4000 {
		fmt.Println("x e 4000")
	} else if x >= 2000 && x <= 3000 {
		fmt.Println("x e maio que 2000 e menor que 3000")
	} else {
		fmt.Println("x e menor que 2000 ou maior que 3000")
	}
}
