package main

import (
	"fmt"
)

func main() {
	primeiroSlice := []int{1, 2, 3, 4, 5}
	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(segundoSlice)  // aqui ele mostra somete o tamanho do slice com o mesmo subjacente so primeiro slice
	fmt.Println(primeiroSlice) // como ele utiliza o mesmo subjacente o primeiro, ele vai modificar o sub do primeiro.
}
