package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{"Go", 30}
	t := template.Must(
		template.New("CursoTemplate").
			Parse("Curso: {{.Nome}} e carga horária: {{.CargaHoraria}} \n"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
