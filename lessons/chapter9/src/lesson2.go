package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 20, 30, 40, 50, 60, 70, 80, 90}

	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Printf("%T\n", slice)
}
