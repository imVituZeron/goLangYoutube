package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	prints(30)
	wg.Wait()
	fmt.Println("Done")
}

func prints(i int) {
	wg.Add(i)
	for j := 1; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Hi, I'm goroutines:", i+1)
			wg.Done()
		}(x)
	}
}
