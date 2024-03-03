package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	//tipo do array
	fmt.Printf("O tipo de meuArray é %T\n", meuArray)

	// todo o array
	fmt.Println(meuArray)

	// tamanho do array
	fmt.Println("Tamanho do array", len(meuArray))

	// ultimo elemento do array
	// meyArray[tamnho-1]
	fmt.Println("Ultimo elemento do array", meuArray[len(meuArray)-1])

	// loop para percorrer um array
	// for:
	// i variavel de inicialização, confição de parada (tamanho -1), incremento de i++
	for i := 0; i < len(meuArray); i++ {
		fmt.Println("Elemento", i, "do array", meuArray[i])
	}

}
