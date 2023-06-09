package main

import "fmt"

func main() {
	channel := make(chan int)
	channel1 := make(chan int, 1) // maneira de executar o codigo com um buffer, sem uma func anonima
	channel1 <- 42

	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel1)
}
