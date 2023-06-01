package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

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
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}

}
