package main

import "fmt"

func main() {

	anonimo := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemCom map[string]string
	}{
		nome:  "feijao",
		sabor: "salgado",
		ondetem: []string{
			"restaurantes", "casa",
		},
		vaibemCom: map[string]string{
			"feijao":  "arroz",
			"feijao1": "carne",
		},
	}

	fmt.Println(anonimo)
}
