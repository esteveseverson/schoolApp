package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func NotaRoutes(r *gin.Engine) {

	// Rota para listar todas as notas
	r.GET("/notas", func(c *gin.Context) {
		rows, err := config.DB.Query(`
			SELECT 
				n.id, 
				n.aluno_id, 
				n.professor_id, 
				a.turma_id, 
				n.atividade_id, 
				a.valor AS valor_total, 
				n.valor_obtido 
			FROM notas n 
			JOIN atividades a ON n.atividade_id = a.id
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var notas []models.Nota
		for rows.Next() {
			var nota models.Nota
			if err := rows.Scan(&nota.ID, &nota.AlunoID, &nota.ProfessorID, &nota.TurmaID, &nota.AtividadeID, &nota.ValorTotal, &nota.ValorObtido); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			notas = append(notas, nota)
		}
		c.JSON(http.StatusOK, notas)
	})

	// Rota para buscar uma nota pelo ID
	r.GET("/notas/:id", func(c *gin.Context) {
		id := c.Param("id")

		var nota models.Nota
		var valorTotal float64
		var turmaID int

		query := `
			SELECT 
				n.id, 
				n.aluno_id, 
				n.professor_id, 
				a.turma_id, 
				n.atividade_id, 
				a.valor AS valor_total, 
				n.valor_obtido 
			FROM notas n 
			JOIN atividades a ON n.atividade_id = a.id
			WHERE n.id = $1
		`
		err := config.DB.QueryRow(query, id).Scan(&nota.ID, &nota.AlunoID, &nota.ProfessorID, &turmaID, &nota.AtividadeID, &valorTotal, &nota.ValorObtido)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		nota.TurmaID = turmaID
		nota.ValorTotal = valorTotal
		c.JSON(http.StatusOK, nota)
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
		var turmaID int
		err = config.DB.QueryRow("SELECT valor, turma_id FROM atividades WHERE id = $1", nota.AtividadeID).Scan(&valorTotal, &turmaID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar atividade"})
			return
		}

		// Verificar se o aluno já possui uma nota para essa atividade
		var notaExistenteID int
		err = config.DB.QueryRow("SELECT id FROM notas WHERE aluno_id = $1 AND atividade_id = $2", nota.AlunoID, nota.AtividadeID).Scan(&notaExistenteID)
		if err != nil && err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar nota existente"})
			return
		}
		if notaExistenteID >= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O aluno já possui uma nota para essa atividade"})
			return
		}

		// Verificar se o valor obtido não excede o valor total da atividade
		if nota.ValorObtido > *&valorTotal {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O valor obtido não pode exceder o valor total da atividade"})
			return
		}

		// Inserir nota no banco de dados
		query := `INSERT INTO notas (aluno_id, atividade_id, professor_id, valor_obtido) VALUES ($1, $2, $3, $4) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, nota.AlunoID, nota.AtividadeID, nota.ProfessorID, nota.ValorObtido).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		nota.ID = id
		nota.TurmaID = turmaID
		nota.ValorTotal = valorTotal
		c.JSON(http.StatusOK, nota)
	})

	// Rota para atualizar uma nota existente
	r.PUT("/notas/:id", func(c *gin.Context) {
		id := c.Param("id")
		var nota models.Nota
		if err := c.ShouldBindJSON(&nota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a atividade existe e obter o valor total
		var valorTotal float64
		err := config.DB.QueryRow("SELECT valor FROM atividades WHERE id = $1", nota.AtividadeID).Scan(&valorTotal)
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
		query := `UPDATE notas SET aluno_id = $1, atividade_id = $2, professor_id = $3, valor_obtido = $4 WHERE id = $5`
		_, err = config.DB.Exec(query, nota.AlunoID, nota.AtividadeID, nota.ProfessorID, nota.ValorObtido, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Nota atualizada com sucesso"})
	})

	// Rota para deletar uma nota existente
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
