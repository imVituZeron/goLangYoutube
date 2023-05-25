package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (x ordenarPorPotencia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (x ordenarPorConsumo) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (x ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	carros := []carro{carro{"Uno", 400, 17},
		carro{"Porsche", 2, 1},
		carro{"Fusca", 500, 30},
	}
	fmt.Println("Inicial:\n", carros)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println("Potencia:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println("Consumo:\n", carros)
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println("Lucro para o Posto:\n", carros)

}
