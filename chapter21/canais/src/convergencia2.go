package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("asdf")

	newNew := converge(work("bebe"), work("lele"))

	for i := 0; i < 15; i++ {
		fmt.Println(<-newNew)
	}

}

func work(s string) chan string {
	channel := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Funcao %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, channel)

	return channel
}

func converge(x, y chan string) chan string {
	newChannel := make(chan string)

	go func() {
		for {
			newChannel <- <-x
		}
	}()

	go func() {
		for {
			newChannel <- <-y
		}
	}()

	return newChannel
}
