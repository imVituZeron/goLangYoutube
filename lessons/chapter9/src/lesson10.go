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

	slice["bemjor_jorge"] = []string{"cantar", "compor"}
	delete(slice, "silva_jorge") // deletando value do map

	for key, value := range slice {
		fmt.Println(key)
		for _, vlue := range value {
			fmt.Println("\t\t", vlue)
		}
	}

}
