package main

import "fmt"

// slice são um tipo de array dinamico, que não tem um tamanho fixo

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(s)

	// o tipo slice ainda é um array
	fmt.Printf("O tipo de s é %T\n", s)

	// tamanho do slice
	fmt.Println("Tamanho do slice", len(s))

	// tamanho e capacidade
	fmt.Printf("tamnaho = %d, capacidade = %d\n", len(s), cap(s))

	// cortando dados para analizar tamanho e capacidade

	// pegar do inicio até a terceira posição
	fmt.Printf("tamnaho = %d, capacidade = %d, array = %d\n", len(s[:3]), cap(s[:3]), s[:3])

	// imprimir da posição 4 até o final
	fmt.Printf("tamnaho = %d, capacidade = %d, array = %d\n", len(s[4:]), cap(s[4:]), s[4:])

	// aumentado o tamno do slice dinamicamente
	//append é uma função que adiciona um elemento ao slice

	s = append(s, 6)
	// note que mesmo adicionando um elemento a capacidade do slice dobrou
	// isso ocorre porque o go cria um novo array com a mesma capacidade do anterior e
	// monta um ponteiro para interligar os dois arrays (foi o que eu entendi)
	// de toda forma, tem que cuidar, porque crescer o slice dinamicamente pode ser custoso
	fmt.Printf("tamanho = %d, capacidade = %d, array = %d\n", len(s), cap(s), s)
}
