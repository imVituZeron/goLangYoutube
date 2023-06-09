package main

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 64
	}()
	fmt.Println(<-cs)

	fmt.Printf("--------\n")
	fmt.Printf("cs\t%T\n", cs)
}
