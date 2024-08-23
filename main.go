package main

import (
	"schoolApp/config"
	"schoolApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Conectar ao banco de dados
	config.ConnectDB()

	// Registrar rotas
	routes.ProfessorRoutes(r)
	routes.TurmaRoutes(r)
	routes.AlunoRoutes(r)
	routes.AtividadeRoutes(r)
	routes.NotaRoutes(r)

	// Iniciar o servidor
	r.Run(":8080")
}
