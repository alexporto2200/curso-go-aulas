package main

import (
	"errors"
	"fmt"
)

func main() {
	// chamando funcao sum
	var t = sum(1, 2)
	fmt.Println(t)

	// chamando funcao sum2
	t = sum2(1, 2)
	fmt.Println(t)

	fmt.Println("-------------------")
	// recebendo multiplos retornos
	fmt.Println(sum3(3, 4))

	// recebendo multiplos retornos
	fmt.Println(sum3(3, 0))

	fmt.Println("-------------------")
	// recebendo multiplos retornos
	// tipo erro
	fmt.Println(sum4(3, 0))

	// recendo os dados em variavel
	fmt.Println("-------------------")
	valor, err := sum4(1, 1)
	if err == nil {
		fmt.Println(valor)
	} else {
		fmt.Println(err)
	}

	// recendo os dados em variavel
	fmt.Println("-------------------")
	valor, err = sum4(4, 1)
	if err == nil {
		fmt.Println(valor)
	} else {
		fmt.Println(err)
	}
}

// dois parametros e um retorno
func sum(a int, b int) int {
	return a + b
}

// outro tipo de declarar os parametros de entrada, quando s~ao do mesmo tipo
func sum2(a, b int) int {
	return a + b
}

// funcao com multiplos retornos
func sum3(a, b int) (int, bool) {
	result := a + b

	if result > 3 {
		return result, true
	} else {
		return result, false
	}
}

func sum4(a, b int) (int, error) {
	result := a + b

	if result > 3 {
		// nil é o o tipo nulo em go
		return result, nil
	} else {
		return result, errors.New("Resultado menor que 3")
	}
}
