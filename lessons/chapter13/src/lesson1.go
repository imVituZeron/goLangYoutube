package main

import "fmt"

func main() {
	fmt.Println(retornaInt(1))
	fmt.Println(retornaIntString(145, "vitor"))
}

func retornaInt(x int) int {
	return x + 1
}

func retornaIntString(x int, nome string) (int, string) {
	return x, nome
}
