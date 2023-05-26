package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1950"

	//transforma a senha em um hash
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sb)
	// compara o hash gerado com a funcao
	// senha bate com o hash, entao vai trazer nada
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))

	// senha NAO bate com o hash, entao vai trazer um erro.
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}
