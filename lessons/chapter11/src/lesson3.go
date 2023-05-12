package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	saveiro := caminhonete{veiculo{2, "preto"}, true}

	unoweekend := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "braco com escada em cima",
		},
		modeloLuxo: false,
	}

	fmt.Println(saveiro)
	fmt.Println(unoweekend)

	fmt.Println(saveiro.cor)
	fmt.Println(unoweekend.modeloLuxo)
}
