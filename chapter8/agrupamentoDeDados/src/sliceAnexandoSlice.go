package main

import (
	"fmt"
)

func main() {
	umaSlice := []int{1, 2, 3, 4}
	outraSlice := []int{}

	outraSlice = append(outraSlice, 9, 10, 11, 12)
	umaSlice = append(umaSlice, 5, 6, 7, 8)
	maisUmaSlice := append(umaSlice, outraSlice...) // os tres pontos diz que vai pegar os items dos slice e nao o slice
	// por default uma slice do tipo int nao aceita um slice.

	fmt.Println(maisUmaSlice)
}
