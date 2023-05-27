package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	var s = []byte(`[{"Name":"Sara","Age":21},{"Name":"Vitor","Age":23},{"Name":"Eliza","Age":42}]`)

	var P []Pessoa

	err := json.Unmarshal(s, &P)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(P)
}
