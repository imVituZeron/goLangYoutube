package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Vitor",
		sobrenome: "Santos",
		fumante:   false,
	}

	cliente2 := cliente{"Sara", "Stefani", false} // outro modo de instanciar uma struct
	fmt.Println(cliente1, cliente2)
}
