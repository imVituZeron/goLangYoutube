package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		[]string{
			"Sara Stefani",
			"Leuterio",
			"Instagram",
		},
		[]string{
			"Charles",
			"augusto",
			"Redes",
		},
		[]string{
			"Jorge",
			"Anfisa",
			"Se bobo",
		},
	}
	fmt.Println(ss)
	for _, value := range ss {
		for _, v := range value {
			fmt.Println(v)
		}

	}
}
