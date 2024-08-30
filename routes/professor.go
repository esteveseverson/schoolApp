package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func ProfessorRoutes(r *gin.Engine) {

	// Rota para listar todos os professores
	r.GET("/professores", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, nome, email, cpf FROM professores")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var professores []models.Professor
		for rows.Next() {
			var professor models.Professor
			if err := rows.Scan(&professor.ID, &professor.Nome, &professor.Email, &professor.CPF); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			professores = append(professores, professor)
		}
		c.JSON(http.StatusOK, professores)
	})

	// Rota para buscar um professor pelo ID
	r.GET("/professores/:id", func(c *gin.Context) {
		id := c.Param("id")

		var professor models.Professor

		query := `SELECT id, nome, email, cpf FROM professores WHERE id = $1`
		err := config.DB.QueryRow(query, id).Scan(&professor.ID, &professor.Nome, &professor.Email, &professor.CPF)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Professor não encontrado"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, professor)
	})

	// Rota para criar um novo professor
	r.POST("/professores", func(c *gin.Context) {
		var professor models.Professor
		if err := c.ShouldBindJSON(&professor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO professores (nome, email, cpf) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err := config.DB.QueryRow(query, professor.Nome, professor.Email, professor.CPF).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		professor.ID = id
		c.JSON(http.StatusOK, professor)
	})

	// Rota para atualizar um professor existente
	r.PUT("/professores/:id", func(c *gin.Context) {
		id := c.Param("id")
		var professor models.Professor
		if err := c.ShouldBindJSON(&professor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `UPDATE professores SET nome=$1, email=$2, cpf=$3 WHERE id=$4`
		_, err := config.DB.Exec(query, professor.Nome, professor.Email, professor.CPF, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Professor atualizado com sucesso"})
	})

	// Rota para deletar um professor existente
	r.DELETE("/professores/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Verificar se o professor está associado a alguma turma
		var turmaNome string
		queryCheck := `
			SELECT t.nome 
			FROM turmas t 
			WHERE t.professor_id = $1
		`
		err := config.DB.QueryRow(queryCheck, id).Scan(&turmaNome)
		if err == nil {
			// Se encontrar uma turma associada, retorna erro
			c.JSON(http.StatusBadRequest, gin.H{"error": "Professor está cadastrado na Turma", "turma": turmaNome})
			return
		} else if err != sql.ErrNoRows {
			// Caso ocorra um erro diferente de "sem linhas", retorna erro interno
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turmas associadas ao professor"})
			return
		}

		// Se o professor não estiver associado a nenhuma turma, procede com a exclusão
		query := `DELETE FROM professores WHERE id=$1`
		_, err = config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Professor excluído com sucesso"})
	})
}
