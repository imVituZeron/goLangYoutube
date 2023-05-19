package main

import (
	"fmt"
)

func main() {
	sliceOfInt := []int{30, 30, 30}
	fmt.Println(retornaSoma(sliceOfInt...))
	fmt.Println(retorna(sliceOfInt))
}

func retornaSoma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func retorna(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
