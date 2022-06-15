package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=root sslmode=disable host=localhost"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

type Produto struct {
	id         int
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
	// produtos := []Produto{
	// 	{Nome: "Camiseta",
	// 		Descricao:  "Azul, bem bonita",
	// 		Preco:      39,
	// 		Quantidade: 10},
	// 	{"Tenis", "Conforatável", 89, 3},
	// 	{"Fone", "Muito bom", 59, 2},
	// }

	db := conectaComBancoDeDados()

	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
