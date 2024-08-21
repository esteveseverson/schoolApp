package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

var DB *sql.DB

func Connect() {
	// docker run --name postgres-schoolapp -e POSTGRES_DB=schoolApp -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123123 -p 5432:5432 -d postgres
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
