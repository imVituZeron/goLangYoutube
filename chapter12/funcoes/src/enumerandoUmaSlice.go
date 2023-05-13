package main

import "fmt"

func main() {

	//sliceInt := []int{1, 2, 3, 4, 5, 6, 17}
	value := soma()
	fmt.Println(value)
}

// quando os parementros da funcao sao variadicos pode ser passado o valor zero.
func soma(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
