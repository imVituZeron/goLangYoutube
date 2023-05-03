package main

import (
	"fmt"
)

func testeInt() {
	sliceInt := []int{1, 2, 4, 5, 6, 7, 8, 10}

	total := 0
	for _, value := range sliceInt {
		total += value

		fmt.Printf("Value %v, soma: %v\n", value, total)
	}
}

func main() {
	slice := []string{"banana", " maca", "jaca"}

	slice = append(slice, "melancia") // adiciona um valor no slice

	for index, value := range slice { // o range ele itera sobre os valores de um slice
		fmt.Printf("No indice %v, temos o valor %v\n", index, value)
	}

	testeInt()
}
