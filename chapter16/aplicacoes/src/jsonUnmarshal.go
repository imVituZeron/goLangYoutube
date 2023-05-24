package main

import (
	"encoding/json"
	"fmt"
)

type Informacoes struct {
	Nome      string  `json:"Nome"`
	Sobrenome string  `json:"Sobrenome"`
	Idade     int     `json:"Idade"`
	Profissao string  `json:"Profissao"`
	Conta     float64 `json:"Conta"`
}

func main() {
	sliceOfBytes := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"espiao","Conta":30}`)

	var info Informacoes
	err := json.Unmarshal(sliceOfBytes, &info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

}
