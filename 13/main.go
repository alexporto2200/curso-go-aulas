package main

func soma(a, b int) int {
	return a + b
}

func somaPonteiros(a, b *int) int {
	return *a + *b
}

func alterarPorReferencia(a *int) {
	*a = 100
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	// soma por copia do valor
	a := soma(minhaVar1, minhaVar2)
	println(a)
	// soma por referencia
	a = somaPonteiros(&minhaVar1, &minhaVar2)
	println(a)

	// alterar por referencia
	alterarPorReferencia(&minhaVar1)
	println(minhaVar1)

}
