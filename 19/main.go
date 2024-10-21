package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	for i, numero := range numeros {
		println(i, numero)
	}

	// blank identifier
	for _, numero := range numeros {
		println(numero)
	}
	// blank identifier (index only)
	for i, _ := range numeros {
		println(i)
	}

	// for tipo um loop

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// for infinito
	for {
		println("loop infinito, use control c para parar")
	}
}
