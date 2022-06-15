package models

import "loja/db"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

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

	defer db.Close()
	return produtos
}
