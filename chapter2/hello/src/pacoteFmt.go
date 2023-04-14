package main

import "fmt"

func main() {
	x := `"ola \ncomo voce vai\n?"` // raw string = não é interpretada (não interpreta os pula linha)
	z := "ola \ncomo voce vai\n?"   // interpreted string = é interpretada (interpreta os pula linha)

	fmt.Println(x)
	fmt.Println(z)
	//=============================

	fmt.Print("oi")
	fmt.Printf("oi")
	fmt.Println("oi") //adiciona uma linha nova no final
	//==============================
	rest := fmt.Sprint("oi", "ADFADF")         // pega as strings e coloca em uma variavel
	result := fmt.Sprintf(x, "oi", "tudo bem") // mesma coisa que a anterios mas com formatação
	response := fmt.Sprintln("oi")             // mesma que as anteriores mas entra e string pulando a linha no final

	fmt.Println(rest)
	fmt.Println(result)
	fmt.Println(response)
}
