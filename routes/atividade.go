package routes

import (
	"net/http"
	"schoolApp/config"
	"schoolApp/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AtividadeRoutes(r *gin.Engine) {
	r.GET("/atividades", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, turma_id, valor, data FROM atividades")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var atividades []models.Atividade
		for rows.Next() {
			var atividade models.Atividade
			if err := rows.Scan(&atividade.ID, &atividade.TurmaID, &atividade.Valor, &atividade.Data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			atividades = append(atividades, atividade)
		}
		c.JSON(http.StatusOK, atividades)
	})

	r.POST("/atividades", func(c *gin.Context) {
		var atividade models.Atividade
		if err := c.ShouldBindJSON(&atividade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a turma existe
		var turmaID int
		err := config.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", atividade.TurmaID).Scan(&turmaID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
			return
		}

		query := `INSERT INTO atividades (turma_id, valor, data) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, atividade.TurmaID, atividade.Valor, time.Now()).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		atividade.ID = id
		c.JSON(http.StatusOK, atividade)
	})
}
