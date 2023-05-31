package main

import "fmt"

type person struct {
	name string
	last string
	age  int
}

func (p *person) speak() {
	fmt.Printf("Hi, i'm %v %v\n", p.name, p.last)
}

type humans interface {
	speak()
}

func talkToMe(h humans) {
	h.speak()
}

func main() {

	p1 := person{"Vitor", "Santos", 45}
	p2 := person{"Sara", "Stefani", 43}

	talkToMe(&p1) // quando voce passa o endereco de p1 voce ta passando os metodos dele tambem.
	talkToMe(&p2)
}
