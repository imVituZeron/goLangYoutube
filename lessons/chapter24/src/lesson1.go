package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Loves []string
}

func main() {
	p1 := person{
		First: "Geraldo",
		Last:  "de Rivia",
		Loves: []string{"Yennefer", "Triss", "Bruxa"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}
