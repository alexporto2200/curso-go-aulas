package main

import (
	"fmt"

	"github.com/alexporto2200/curso-go-aulas/18-1/matematica"
	"github.com/google/uuid"
)

/*
	Iniciando o modulo
	go mod init '18-1'
	go mod tidy
*/

func main() {
	s := matematica.Soma(3, 2)
	fmt.Println("O resultado é ", s)
	fmt.Println("O valor de A é ", matematica.A)

	fmt.Println("O valor de UUID é ", uuid.New())
}
