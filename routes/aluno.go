package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
)

func AlunoRoutes(r *gin.Engine) {
	// Rota para listar todos os alunos
	r.GET("/alunos", func(c *gin.Context) {
		rows, err := config.DB.Query(`
			SELECT a.id, a.nome, a.matricula, at.turma_id
			FROM alunos a
			LEFT JOIN alunos_turmas at ON a.id = at.aluno_id
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		alunosMap := make(map[int]*models.Aluno)
		for rows.Next() {
			var id int
			var nome string
			var matricula string
			var turmaID sql.NullInt32

			if err := rows.Scan(&id, &nome, &matricula, &turmaID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			aluno, exists := alunosMap[id]
			if !exists {
				aluno = &models.Aluno{
					ID:        id,
					Nome:      nome,
					Matricula: matricula,
					TurmaIDs:  []int{},
				}
				alunosMap[id] = aluno
			}

			if turmaID.Valid {
				aluno.TurmaIDs = append(aluno.TurmaIDs, int(turmaID.Int32))
			}
		}

		var alunos []models.Aluno
		for _, aluno := range alunosMap {
			alunos = append(alunos, *aluno)
		}

		c.JSON(http.StatusOK, alunos)
	})

	// Rota para criar um novo aluno
	r.POST("/alunos", func(c *gin.Context) {
		var aluno models.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se cada turma existe
		for _, turmaID := range aluno.TurmaIDs {
			var id int
			err := config.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", turmaID).Scan(&id)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
				return
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turma"})
				return
			}
		}

		// Inserir o aluno
		query := `INSERT INTO alunos (nome, matricula) VALUES ($1, $2) RETURNING id`
		var id int
		err := config.DB.QueryRow(query, aluno.Nome, aluno.Matricula).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Associar o aluno às turmas
		for _, turmaID := range aluno.TurmaIDs {
			_, err = config.DB.Exec("INSERT INTO alunos_turmas (aluno_id, turma_id) VALUES ($1, $2)", id, turmaID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		aluno.ID = id
		c.JSON(http.StatusOK, aluno)
	})

	// Rota para atualizar um aluno existente
	r.PUT("/alunos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var aluno models.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se cada turma existe
		for _, turmaID := range aluno.TurmaIDs {
			var tid int
			err := config.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", turmaID).Scan(&tid)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
				return
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turma"})
				return
			}
		}

		// Atualizar o aluno
		query := `UPDATE alunos SET nome = $1, matricula = $2 WHERE id = $3`
		_, err := config.DB.Exec(query, aluno.Nome, aluno.Matricula, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Limpar as turmas existentes do aluno na tabela aluno_turmas
		_, err = config.DB.Exec("DELETE FROM alunos_turmas WHERE aluno_id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao limpar turmas do aluno"})
			return
		}

		// Associar o aluno às novas turmas
		for _, turmaID := range aluno.TurmaIDs {
			_, err = config.DB.Exec("INSERT INTO alunos_turmas (aluno_id, turma_id) VALUES ($1, $2)", id, turmaID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao associar aluno às turmas"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Aluno atualizado com sucesso"})
	})

	// Rota para deletar um aluno existente
	r.DELETE("/alunos/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := `DELETE FROM alunos WHERE id = $1`
		_, err := config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado com sucesso"})
	})
}
