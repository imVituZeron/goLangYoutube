package main

import (
	"fmt"
)

func main() {
	esportaFavorito := "Bocha"

	switch esportaFavorito {
	case "Basquete":
		fmt.Println("Sara gosta de Basquete")
	case "Futebol":
		fmt.Println("Sara gosta de Futebol")
	case "Atletismo":
		fmt.Println("Sara gosta de Atletismo")
	case "Bocha":
		fmt.Println("Sara gosta de Bocha")
	default:
		fmt.Println("Sara gosta de gosta de todos!")
	}

}
