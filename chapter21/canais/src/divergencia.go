package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go manda(20, c1)
	go outra(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(x int) {
			c2 <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))

	return n
}
