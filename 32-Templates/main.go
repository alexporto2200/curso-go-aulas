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
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} e carga horária: {{.CargaHoraria}} \n")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
