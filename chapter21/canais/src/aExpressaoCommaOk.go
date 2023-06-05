package main

import "fmt"

func main() {
	chann := make(chan int)

	go func() {
		chann <- 42
		close(chann)
	}()

	v, ok := <-chann
	fmt.Println(v, ok)

	v, ok = <-chann
	fmt.Println(v, ok)
}
