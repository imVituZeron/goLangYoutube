package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Name string
	Age  int
}

func main() {
	u1 := usuario{"Sara", 21}
	u2 := usuario{"Vitor", 23}
	u3 := usuario{"Eliza", 42}

	users := []usuario{u1, u2, u3}
	fmt.Println("structs: ", users)

	b, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("jsons: ", string(b))

}
