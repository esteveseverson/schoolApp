package routes

import (
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func ProfessorRoutes(r *gin.Engine) {
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

	r.DELETE("/professores/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := `DELETE FROM professores WHERE id=$1`
		_, err := config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Professor excluído com sucesso"})
	})
}
