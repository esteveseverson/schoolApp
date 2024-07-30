package routes

import (
	"sistema-web/database"
	"sistema-web/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/professores", func(c *gin.Context) {
		var professor models.Professor
		if err := c.ShouldBindJSON(&professor); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&professor)
		c.JSON(200, professor)
	})

	// Adicione rotas para Turmas, Alunos, Atividades e Notas
}
