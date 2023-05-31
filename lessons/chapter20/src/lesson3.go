package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const GoroutinesCount = 500

func main() {
	createGoRoutines(GoroutinesCount)
	wg.Wait()
	fmt.Println("total de goroutines:", GoroutinesCount, "\nTotal do contador:", contador)
}

func createGoRoutines(i int) {
	wg.Add(i)
	for y := 0; y < i; y++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}

}
