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

func (c Cliente) Desativar() {
	c.ativo = false
	fmt.Printf("Cliente %s desativado\n", c.nome)
}

func main() {

	alex := Cliente{
		nome:  "Alex",
		idade: 30,
		ativo: true,
		Endereco: Endereco{
			rua:    "Rua 1",
			numero: 123,
			cidade: "São Paulo",
			estado: "SP",
		}}
	alex.Desativar()

	fmt.Println(alex)
}
