package main

import "fmt"

type hotdog int // criando um tipo novo

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("tipos da x: %T\n", x)
	fmt.Printf("tipos da b: %T\n", b) // as duas tem tipos diferentes

	fmt.Printf("valor da b: %v\n", b)
	fmt.Printf("valor da x: %v\n", x)
	x = int(b) // convertendo b e int para o x receber o seu valor
	fmt.Printf("valor de x depois do conversão: %v\n", x)

	fmt.Println("type is life")
}
