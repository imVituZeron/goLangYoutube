package main

import "fmt"

func main() {
	channel := make(chan int)
	go coloca(channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func coloca(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
