package main

import "fmt"

type erroEspecial struct {
	anyThing string
}

func (e erroEspecial) Error() string {
	return "deu zica"
}

func main() {
	arg := erroEspecial{"EMOJIS"}
	erroComoParamentro(arg)
}

func erroComoParamentro(e error) {
	fmt.Println("passaram como argumento para mim o seguinte: ", e.(erroEspecial).anyThing)
}
