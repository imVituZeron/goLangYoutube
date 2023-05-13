package main

import "fmt"

func main() {
	basica()

	soma(10, 10)

	value := soma1(1, 1)
	fmt.Println(value)

	total, elementos := soma2(2, 3, 4, 5, 6, 7)
	fmt.Println(total, elementos)

}

// funcao basica
func basica() {
	fmt.Println("Hello, GoodMorning!")
}

// funcao que aceita argumento
func soma(x, y int) {
	fmt.Printf("%v + %v = %v\n", x, y, x+y)
}

// funcao com return
func soma1(x, y int) int {
	return x + y
}

// funcao com varios retornos e com paramentos variadico
func soma2(x ...int) (int, int) {
	resultado := 0
	for _, v := range x {
		resultado += v
	}

	return resultado, len(x)
}
