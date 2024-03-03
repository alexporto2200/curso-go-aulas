package main

type Cliente struct {
	nome string
}

// altera o valor do nome
// pelo ponteiro
func (c *Cliente) Andou() {
	c.nome = "João"
	println(c.nome, "andou")
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) Depositar(valor int) {
	c.saldo += valor
}

func main() {
	alex := Cliente{
		nome: "Alex",
	}

	alex.Andou()

	println(alex.nome)

	println("------------")

	conta := NewConta()
	println(conta.saldo)

	conta.Depositar(100)
	println(conta.saldo)
}
