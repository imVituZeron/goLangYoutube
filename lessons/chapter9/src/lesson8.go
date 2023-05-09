package main

import (
	"fmt"
)

func main() {
	slice := map[string][]string{
		"silva_jorge": []string{
			"comprar",
			"carros",
		},
		"Strainger_eleven": []string{
			"desaparecer",
			"raspar o cabelo",
		},
		"pike_rob": []string{
			"codar",
			"criar liguagens",
		},
	}
	for key, value := range slice {
		fmt.Println(key)
		for _, vlue := range value {
			fmt.Println("\t\t", vlue)
		}
	}
	fmt.Println(slice)
}
