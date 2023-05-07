package main

import (
	"fmt"
)

func main() {

	slice := []int{11, 22, 30, 40, 54, 66, 77, 88, 99, 1010}

	slicePrimeiro := append(slice[:3])
	sliceSegundo := append(slice[4:])
	sliceTerceiro := append(slice[1:7])
	sliceQuarto := append(slice[2 : len(slice)-1])
	fmt.Println(slicePrimeiro)
	fmt.Println(sliceSegundo)
	fmt.Println(sliceTerceiro)
	fmt.Println(sliceQuarto)
}
