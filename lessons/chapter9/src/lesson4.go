package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)    // <-- demonstrando o slice normal
	x = append(x, 52) // <-- add um value
	fmt.Println(x)
	x = append(x, 53, 54, 55) // <-- add varios value
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // <-- add o slice y o slice x
	fmt.Println(x)
}
