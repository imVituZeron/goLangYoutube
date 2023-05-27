package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type usuario struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	u1 := usuario{"Sara", 21}
	u2 := usuario{"Vitor", 23}
	u3 := usuario{"Eliza", 42}

	users := []usuario{u1, u2, u3}
	fmt.Println("structs: ", users)

	err := json.NewEncoder(os.Stdout).Encode(users) // usando methoed chaning

	if err != nil {
		fmt.Println("aqui ja era", err)
	}

}
