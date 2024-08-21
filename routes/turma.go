package routes

import (
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func TurmaRoutes(r *gin.Engine) {
	r.GET("/turmas", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, nome, ano, professor_id, semestre FROM turmas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var turmas []models.Turma
		for rows.Next() {
			var turma models.Turma
			if err := rows.Scan(&turma.ID, &turma.Nome, &turma.Ano, &turma.ProfessorID, &turma.Semestre); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			turmas = append(turmas, turma)
		}
		c.JSON(http.StatusOK, turmas)
	})

	r.POST("/turmas", func(c *gin.Context) {
		var turma models.Turma
		if err := c.ShouldBindJSON(&turma); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se o professor existe
		var professorID int
		err := config.DB.QueryRow("SELECT id FROM professores WHERE id = $1", turma.ProfessorID).Scan(&professorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Professor inválido"})
			return
		}

		query := `INSERT INTO turmas (nome, ano, professor_id, semestre) VALUES ($1, $2, $3, $4) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, turma.Nome, turma.Ano, turma.ProfessorID, turma.Semestre).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		turma.ID = id
		c.JSON(http.StatusOK, turma)
	})
}
