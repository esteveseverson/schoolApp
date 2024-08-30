package main

import (
	"schoolApp/config"
	"schoolApp/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar o router
	r := gin.Default()

	// Configurar o middleware de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://167.234.254.164:8080"},
		//AllowOrigins:     []string{"http://localhost:8088"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Conectar ao banco de dados
	config.ConnectDB()

	// Registrar rotas
	routes.ProfessorRoutes(r)
	routes.TurmaRoutes(r)
	routes.AlunoRoutes(r)
	routes.AtividadeRoutes(r)
	routes.NotaRoutes(r)

	// Iniciar o servidor
	r.Run(":8088")
}
