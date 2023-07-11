package main

import testing

type test struct{
	data []int
	anwser int
}

func TestsomaEmTabela(t *testing.T){
	tests := []test{
		test{data: []int{1,2,3}, anwser: 6},
		test{data: []int{10,11,12}, anwser: 33},
		test{data: []int{5,10,-2}, anwser: 10},
	}
	for _, v := range tests {
		z := soma(v.data...)
		if z != v.anwser {
			t.Error("Expected:", v.anwser "Got:", z)
		}
	}
}

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got", teste)
	}
}

func TestMulti(t *testing.T) {
	teste := multi(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got", teste)
	}
}
