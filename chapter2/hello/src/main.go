package main

import "fmt"

func main() {
	numeroBytes, erros := fmt.Println("Hello World!", "oi", 100)
	fmt.Println(numeroBytes, erros)
}
