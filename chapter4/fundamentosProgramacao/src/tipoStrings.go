package main

import "fmt"

func main() {
	str := "Hello Playground"

	strLouca := `Hello 
		playground
							louco!` // usar a crase para criar strings assim <-

	fmt.Println(str)
	fmt.Printf("valor: %v, tipo: %T\n", str, str) // monstrando o valor e o tipo da variavel

	fmt.Println(strLouca)

	strSlice := []byte(str) // criando uma slice de bytes fazendo o conversao de string para bytes
	fmt.Println(strSlice)
	for _, v := range strSlice {
		fmt.Printf("valor: %v, tipo: %T, unicode: %#U, hexa: %#x\n", v, v, v, v)
	}
}
