package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go myloop(10, c)
	myloop2(c)

}

func myloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) // serve para fechar o canal quando ele nao vai receber mais valores
}

func myloop2(s <-chan int) {
	for v := range s {
		fmt.Println("Recebido do canal:", v)
	}
}
