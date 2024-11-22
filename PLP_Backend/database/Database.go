package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaDB() *sql.DB {
	conexao := "user=postgres dbname=TheBoys password=davi25310 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}
	return db
}