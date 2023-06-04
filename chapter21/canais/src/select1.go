package main

import "fmt"

func main() {
	c1 := make(chan int)
	quit := make(chan int)
	go func1(c1, quit)
	func2(c1, quit)
}

func func1(c, q chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("recebido:", <-c)
	}
	q <- 0
}

func func2(c, q chan int) {
	qq := 10
	for {
		select {
		case c <- qq:
			qq++
		case <-q:
			return
		}
	}
}
