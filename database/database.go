package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

var DB *sql.DB

func Connect() {
	connStr := "user=postgres password=123123 dbname=schoolApp sslmode=disable host=localhost port=5432" // Atualize conforme suas configurações
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL!")
}
