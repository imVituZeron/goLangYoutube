package main

import "fmt"

func main() {
	essaVaiRedecerAoutraFuncao(issoVaiSerUmArgumento)
}

func issoVaiSerUmArgumento() {
	fmt.Println("Olha eu aqui")
}

func essaVaiRedecerAoutraFuncao(x func()) {
	fmt.Println("Prestenção")
	x()
}
