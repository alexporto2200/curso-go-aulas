package main

import "fmt"

func main() {
	fmt.Println(sum(10, 11, 12, 123, 13, 100))
}

// funcao variadica aceita mais de um parametro
func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
