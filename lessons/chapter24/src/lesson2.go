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
		Last:  "De rivia",
		Loves: []string{
			"Yennefer", "Triss", "Keira",
		},
	}
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Bino!")
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("caiu na cilada")
	}

	return bs, nil
}
