package database

import (
	"database/sql"
	"log"
)

func ConectaDB() *sql.DB {
	conexao := "user=postgres dbname=TheBoys password=davi252310 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
