package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaCombancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=my_password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
