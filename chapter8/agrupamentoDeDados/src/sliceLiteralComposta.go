package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 7) //mudando o tamanho do slice com a funcao append
	fmt.Println(slice2)

	fmt.Println(slice2[3])
	slice2[3] = 1234 // trocando o valor do slice na posicao 3
	fmt.Println(slice2[3])
	fmt.Println(slice2)
}
