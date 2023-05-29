package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totalDeGoroutines := 100

	wg.Add(totalDeGoroutines)

	var mutex sync.Mutex

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			mutex.Lock() //<-- ele tranca a execucao
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mutex.Unlock() //<-- ele destranca a execucao
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
