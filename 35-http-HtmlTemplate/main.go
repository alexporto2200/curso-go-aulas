package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", ExecTemplate)
	http.ListenAndServe(":8080", nil)
}

func ExecTemplate(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(w, Cursos{
		{"Go", 40},
		{"Java", 200},
		{"php", 10},
	})

	if err != nil {
		panic(err)
	}
}
