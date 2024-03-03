package main

//  globais

// constantes
const a = "Hello World"

// variaveis

// inferencia de dados
var (
	b bool
	c int
	d string
	e float64
	f complex128
)

// inferencia de tipo
var g = 42

//

func main() {
	println("Inferencia de tipos")
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

	// variaveis de escopo da funcao
	var h = 43
	println(h)

	// variaveis short declaration
	j := 44
	println(j)

	// atribuicao de variaveis
	j = 45

	println(j)
}
