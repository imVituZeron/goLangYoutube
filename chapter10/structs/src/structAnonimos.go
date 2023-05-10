package main

import (
	"fmt"
)

func main() {
	x := struct { // forma de usar uma strucs sem declara um tipo novo
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 13,
	}

	fmt.Println(x)
}
