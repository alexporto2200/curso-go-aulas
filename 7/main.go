package main

import "fmt"

func main() {
	salarios := map[string]int{"jose": 1000, "maria": 5000, "pedro": 3000}

	// imprimir todo o map
	fmt.Println(salarios)

	// imprimir um valor do map
	fmt.Println(salarios["jose"])
	fmt.Println("-------------------")
	// remover um valor do map
	delete(salarios, "jose")
	fmt.Println(salarios)
	fmt.Println("-------------------")
	// adicionar um valor ao map
	salarios["jose"] = 3000
	fmt.Println(salarios)
	fmt.Println("-------------------")
	// outra forma de criar um map
	salarios2 := make(map[string]int)
	fmt.Println(salarios2)
	fmt.Println("-------------------")
	// loop para percorrer um map
	for k, v := range salarios {
		fmt.Println(k, v)
	}
	fmt.Println("-------------------")
	// com printf
	for key, value := range salarios {
		fmt.Printf("Nome: %s, Salario: %d\n", key, value)
	}

	fmt.Println("-------------------")
	// ignorar a chave, utilizando o blank identifier
	for _, value := range salarios {
		fmt.Printf("Salario: %d\n", value)
	}

}
