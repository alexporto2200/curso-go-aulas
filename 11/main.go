package main

import "fmt"

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func main() {
	alex := Cliente{"Alex", 28, true}

	fmt.Println(alex)

	// acessando um campo

	fmt.Println("-------------------")

	fmt.Printf("Nome: %s\n", alex.nome)

	fmt.Println("-------------------")

	// alterando um campo
	alex.idade = 29
	fmt.Println(alex)
}
