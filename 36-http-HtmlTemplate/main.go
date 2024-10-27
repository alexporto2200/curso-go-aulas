package main

import (
	"html/template"
	"net/http"
	"strings"
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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
	})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(w, Cursos{
		{"Go", 40},
		{"Java", 200},
		{"php", 10},
	})

	if err != nil {
		panic(err)
	}
}
