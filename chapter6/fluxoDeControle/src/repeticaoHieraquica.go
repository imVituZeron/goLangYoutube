package main

import (
	"fmt"
)

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Printf("hr: %v\n", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Print("\n")
	}
}
