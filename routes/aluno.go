package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func AlunoRoutes(r *gin.Engine) {
	r.GET("/alunos", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, nome, matricula, turma_id FROM alunos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var alunos []models.Aluno
		for rows.Next() {
			var aluno models.Aluno
			if err := rows.Scan(&aluno.ID, &aluno.Nome, &aluno.Matricula, &aluno.TurmaID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			alunos = append(alunos, aluno)
		}
		c.JSON(http.StatusOK, alunos)
	})

	r.POST("/alunos", func(c *gin.Context) {
		var aluno models.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a turma existe
		var turmaID int
		err := config.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", aluno.TurmaID).Scan(&turmaID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turma"})
			return
		}

		query := `INSERT INTO alunos (nome, matricula, turma_id) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, aluno.Nome, aluno.Matricula, aluno.TurmaID).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		aluno.ID = id
		c.JSON(http.StatusOK, aluno)
	})
}
