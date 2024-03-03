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

// interface pessoa
// contrato para metodos
type Pessoa interface {
	Desativar()
}

func desativacao(p Pessoa) {
	p.Desativar()
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

	// alex mesmo sendo do tipo Cliente
	// pode ser passado como parametro
	// para a função desativacao
	// pois Cliente implementa a interface Pessoa
	// e a função desativacao recebe um parametro
	// do tipo Pessoa
	desativacao(alex)

	fmt.Println(alex)
}
