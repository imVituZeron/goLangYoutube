package main

import "fmt"

func main() {
	//x := 5
	//switch {
	//case x < 5:
	//	fmt.Println("x é menor que 5")
	//case x == 5:
	//	fmt.Println("x é igual a 5")
	//case x > 5:
	//	fmt.Println("x é maior que 5")
	//}

	quemTaNoEscritorioHoje := "Zé Geraldo!"

	//switch quemTaNoEscritorioHoje {
	//case "Zé Geraldo!":
	//	fmt.Println("É o Zé")
	//case "Roberto!":
	//	fmt.Println("seÉ o Roberto")
	//case "Charlão!":
	//	fmt.Println("É o Charlão")
	//case "Nathan mais cedo!":
	//	fmt.Println("É o Nathan mais cedo")
	//}

	//switch quemTaNoEscritorioHoje {
	//case "Zé Geraldo!":
	//	fmt.Println("É o Zé")
	//	fallthrough // ele pula aproxima condição e executa o escopo dela !!
	//case "Roberto!":
	//	fmt.Println("sempre que zé vem o Roberto vem também")
	//case "Charlão!":
	//	fmt.Println("É o Charlão")
	//case "Nathan mais cedo!":
	//	fmt.Println("É o Nathan mais cedo")
	//default: // uma resposta padrão para caso não seja true
	//	fmt.Println("Ninguem veio hoje")
	//}

	switch quemTaNoEscritorioHoje {
	case "Zé Geraldo!", "Roberto!": // case com varios valores! pode colcar varios
		fmt.Println("Alguem do time 1 veio!")
	case "Maria!", "Jão!":
		fmt.Println("Alguem do time 2 veio!")
	}

}
