package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectComBancoDeDados() *sql.DB {
	conexao := "user=root dbname=juliano_loja password=soder1989 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
