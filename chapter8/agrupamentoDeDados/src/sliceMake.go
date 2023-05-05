package main

import (
	"fmt"
)

func main() {

	sliceMake := make([]int, 10, 50)

	fmt.Println(sliceMake)
	for x := 0; x < 50; x++ {
		sliceMake = append(sliceMake, x)
	}
	fmt.Println(sliceMake)
}
