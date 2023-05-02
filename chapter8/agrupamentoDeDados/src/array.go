package main

import "fmt"

var x [5]int

// quando voce passa a quantidade de elementos para um array mas utiliza ele zera todos os restante
// igual no resultado da linha 13

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)

	fmt.Println(len(x)) // comprimento de um array
}
