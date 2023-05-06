package main

import "fmt"

func main() {
	qualquerCoisa := map[int]string{
		123: "e muito legal",
		198: "Menos legal",
		983: "Esse e massa",
		118: "Maior de ja",
	}
	fmt.Println(qualquerCoisa)

	for key, value := range qualquerCoisa {
		fmt.Println(key, value) // os maps sao sem  ordem podem aparecer variados nos ranges
	}

	delete(qualquerCoisa, 198) // deletando um item do map

	fmt.Println(qualquerCoisa)
}
