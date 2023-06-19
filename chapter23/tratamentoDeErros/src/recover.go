package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from F")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling G")
	g(0)
	fmt.Println("returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in G", i)
	g(i + 1)
}
