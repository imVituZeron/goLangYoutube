package main

import "fmt"

func main() {
	x := 103

	func(x int) {
		fmt.Println(x, "vezes", 102)
		fmt.Println(x * 102)
	}(x) // <- esse X é como ja tivesse chamando a função
}
