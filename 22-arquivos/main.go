package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tam, err := f.WriteString("AAAAAre")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Foi escrito no arquivo %d bytes\n", tam)

	// escrever bytes
	tamB, err := f.Write([]byte("222222223333333333344444444"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Foi escrito no arquivo %d bytes\n", tamB)

	f.Close()

	// ler arquivo inteiro
	file, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// ler arquivo stream

	fmt.Println("\nLendo arquivo em partes com buffer")

	file2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)

	// ler de 10 em 10 bytes
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			// EOF error
			break
		}

		// imprimir o buffer, n é a posição
		fmt.Println(string(buffer[:n]))
	}

	// remover o arquivo

	os.Remove("arquivo.txt")

}
