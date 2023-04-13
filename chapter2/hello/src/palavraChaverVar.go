package main

import "fmt"

var y = 16 // variavel esta com o escopa do package

func main() {
	y := 10 // variavel esta restrita ao escopo da main
	qualquerCoisa(y)
}

func qualquerCoisa(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
