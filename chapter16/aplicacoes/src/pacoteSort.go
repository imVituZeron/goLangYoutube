package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"Pera", "Murucuja", "Morango", "maca", "gato"}
	si := []int{34, 56, 78, 90, 12, 34}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)
}
