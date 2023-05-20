package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	areaC := math.Pi * 2 * c.raio
	fmt.Printf("A área do circulo é %v\n", areaC)
}

func (q quadrado) area() {
	areaQ := q.lado * q.lado
	fmt.Printf("A área do quadrado é %v\n", areaQ)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {
	x := quadrado{15.0}
	y := circulo{
		raio: 5.25,
	}

	medida(x)
	medida(y)
}
