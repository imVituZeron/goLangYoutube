package main

import "fmt"

func main() {
	retorno := retornFunc()
	retorno()
}

func retornFunc() func() {
	return func() {
		fmt.Println("funcao da funcao!")
	}
}
