package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=root sslmode=disable host=localhost"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
