package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta",
			Descricao:  "Azul, bem bonita",
			Preco:      39,
			Quantidade: 10},
		{"Tenis", "Conforatável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
