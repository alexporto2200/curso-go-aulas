package main

import "fmt"

// variaveis atribuidas
var (
	a string  = "Hello World"
	b bool    = true
	c int     = 42
	d float64 = 3.14
)

// declaracao de tipo
type id int

// atribuicao de variavel com tipo declarado
var f id = 43

func main() {
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O tipo de b é %T\n", b)
	fmt.Printf("O tipo de c é %T\n", c)
	fmt.Printf("O tipo de d é %T\n", d)

	// o tipo personalizado pertence ao pacote main
	fmt.Printf("O tipo de f é %T\n", f)
}
