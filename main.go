package main

import (
	"log"
	"schoolApp/database"
	"schoolApp/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Inicializa a conexão com o banco de dados PostgreSQL
	database.Connect()

	// Configura o router Gin
	router := gin.Default()

	// Configura as rotas
	routes.SetupRoutes(router)

	// Inicia o servidor na porta 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
