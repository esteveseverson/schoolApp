package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func TurmaRoutes(r *gin.Engine) {

	// Rota para listar todas as turmas
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

	// Rota para buscar uma turma pelo ID
	r.GET("/turmas/:id", func(c *gin.Context) {
		id := c.Param("id")

		var turma models.Turma

		query := `SELECT id, nome, ano, professor_id, semestre FROM turmas WHERE id = $1`
		err := config.DB.QueryRow(query, id).Scan(&turma.ID, &turma.Nome, &turma.Ano, &turma.ProfessorID, &turma.Semestre)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, turma)
	})

	// Rota para criar uma nova turma
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

	// Rota para atualizar uma turma existente
	r.PUT("/turmas/:id", func(c *gin.Context) {
		id := c.Param("id")
		var turma models.Turma
		if err := c.ShouldBindJSON(&turma); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `UPDATE turmas SET nome=$1, ano=$2, professor_id=$3, semestre=$4 WHERE id=$5`
		_, err := config.DB.Exec(query, turma.Nome, turma.Ano, turma.ProfessorID, turma.Semestre, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Turma atualizada com sucesso"})
	})

	// Rota para deletar uma turma existente
	r.DELETE("/turmas/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := `DELETE FROM turmas WHERE id=$1`
		_, err := config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Turma excluída com sucesso"})
	})
}
