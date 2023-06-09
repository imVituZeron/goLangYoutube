package main

import "fmt"

func main() {
	c := gen()
	receive(c)
}

func gen() <-chan int {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	return channel
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
