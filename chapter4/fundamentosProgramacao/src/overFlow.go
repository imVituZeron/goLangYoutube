package main

import "fmt"

func main() {
	var i uint16
	i = 65535 // limite de uma variavel com o tipo uint16
	fmt.Println("variavel i no limite: ", i)
	i += 1 // quando adiciondo mais 1 ele zera e começa do 0
	fmt.Println("variavel com o limite zerado: ", i)
	i += 1 // começa a adicionar de novo desdo 0
	fmt.Println("variavel com atribuindo desdo começo: ", i)
}
