package main

import (
	"fmt"
	"sync"
)

func main() {

	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go send(par, impar)
	go receive(par, impar, converge)

	for v := range converge {
		fmt.Println("valor:", v)
	}
}

func send(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func receive(p, i, converge chan int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for v := range p {
			converge <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range i {
			converge <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(converge)
}
