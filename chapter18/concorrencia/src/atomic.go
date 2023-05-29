package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUS:", runtime.NumCPU())

	var contador int64
	contador = 0
	totalDeGoroutines := 15

	wg.Add(totalDeGoroutines)

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final:", contador)
}
