package main

import (
	"fmt"
)

func main() {
	x := 300
	y := x << 1

	fmt.Printf("decimal: %d, binario: %b, hexa: %#x\n", x, x, x)
	fmt.Printf("decimal: %d, binario: %b, hexa: %#x\n", y, y, y)
}
