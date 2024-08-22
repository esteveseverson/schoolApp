package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func NotaRoutes(r *gin.Engine) {
	r.GET("/notas", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, aluno_id, atividade_id, professor_id, valor_total, valor_obtido FROM notas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var notas []models.Nota
		for rows.Next() {
			var nota models.Nota
			if err := rows.Scan(&nota.ID, &nota.AlunoID, &nota.AtividadeID, &nota.ProfessorID, &nota.ValorTotal, &nota.ValorObtido); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			notas = append(notas, nota)
		}
		c.JSON(http.StatusOK, notas)
	})

	r.POST("/notas", func(c *gin.Context) {
		var nota models.Nota
		if err := c.ShouldBindJSON(&nota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se o professor existe
		var professorID int
		err := config.DB.QueryRow("SELECT id FROM professores WHERE id = $1", nota.ProfessorID).Scan(&professorID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Professor inválido"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar professor"})
			return
		}

		// Verificar se a atividade existe e obter o valor total
		var valorTotal float64
		query := `SELECT valor FROM atividades WHERE id = $1`
		err = config.DB.QueryRow(query, nota.AtividadeID).Scan(&valorTotal)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar atividade"})
			return
		}

		// Verificar se o valor obtido não excede o valor total da atividade
		if nota.ValorObtido > valorTotal {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O valor obtido não pode exceder o valor total da atividade"})
			return
		}

		// Inserir nota no banco de dados
		query = `INSERT INTO notas (aluno_id, atividade_id, professor_id, valor_total, valor_obtido) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, nota.AlunoID, nota.AtividadeID, nota.ProfessorID, valorTotal, nota.ValorObtido).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		nota.ID = id
		nota.ValorTotal = valorTotal
		c.JSON(http.StatusOK, nota)
	})

	r.PUT("/notas/:id", func(c *gin.Context) {
		id := c.Param("id")
		var nota models.Nota
		if err := c.ShouldBindJSON(&nota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a atividade existe e obter o valor total
		var valorTotal float64
		query := `SELECT valor FROM atividades WHERE id = $1`
		err := config.DB.QueryRow(query, nota.AtividadeID).Scan(&valorTotal)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar atividade"})
			return
		}

		// Verificar se o valor obtido não excede o valor total da atividade
		if nota.ValorObtido > valorTotal {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O valor obtido não pode exceder o valor total da atividade"})
			return
		}

		// Atualizar nota no banco de dados
		query = `UPDATE notas SET aluno_id = $1, atividade_id = $2, professor_id = $3, valor_total = $4, valor_obtido = $5 WHERE id = $6`
		_, err = config.DB.Exec(query, nota.AlunoID, nota.AtividadeID, nota.ProfessorID, valorTotal, nota.ValorObtido, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Nota atualizada com sucesso"})
	})

	r.DELETE("/notas/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Deletar nota do banco de dados
		query := `DELETE FROM notas WHERE id = $1`
		_, err := config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Nota deletada com sucesso"})
	})
}
