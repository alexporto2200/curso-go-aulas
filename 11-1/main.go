package main

import "fmt"

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Cliente struct {
	nome  string
	idade int
	ativo bool
	// composição de uma struct com outra struct
	Endereco
}

type Cliente_b struct {
	nome  string
	idade int
	ativo bool
	// criando um atributo do tipo Endereco
	endereco Endereco
}

func main() {
	alex := Cliente{"Alex", 28, true,
		Endereco{"Rua 1", 123, "São Paulo", "SP"}}

	fmt.Println(alex)

	fmt.Println("-------------------")

	fmt.Printf("Nome: %s\n", alex.nome)

	fmt.Println("-------------------")

	alex.idade = 29
	fmt.Println(alex)
}
