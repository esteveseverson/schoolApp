package main

import (
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
	database.Connect()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
