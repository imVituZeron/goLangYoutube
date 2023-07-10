package main

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
