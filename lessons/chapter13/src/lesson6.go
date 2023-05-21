package main

import (
	"fmt"
)

func main() {
	x := 6

	func(x int) {
		fmt.Println(x + 10)
	}(x)

}
