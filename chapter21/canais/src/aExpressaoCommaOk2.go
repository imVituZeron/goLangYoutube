package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go send(par, impar, quit)
	receive(par, impar, quit)
}

func send(par, impar chan int, q chan bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	q <- true
}

func receive(par, impar chan int, q chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("recebi do par:", v)
		case v := <-impar:
			fmt.Println("recebi do impar:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("Deu zebra:", v)
			} else {
				fmt.Println("Encerrando! Recebemos: ", v)
			}
			return
		}
	}
}
