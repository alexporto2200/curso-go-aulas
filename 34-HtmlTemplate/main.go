package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(
		template.New("template.html").
			ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 25},
		{"Python", 30},
		{"Java", 40},
	})

	if err != nil {
		panic(err)
	}
}
