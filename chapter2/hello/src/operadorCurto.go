package main

import "fmt"

var mostra = "mostra!" // essa variavel e visivel para todo o codigo
// mostra := "mostra!" -->> não é possivel usar o operador marmota fora de escopo!

func main() {

	x := 10 // essa é visivel somente dentro do espoco da func
	y := "Ola bom dia!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Println(mostra)
}
