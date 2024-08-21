package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/database"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

func SetupRoutes(r *gin.Engine) {
	// Professores
	r.POST("/professores", func(c *gin.Context) {
		var professor models.Professor
		if err := c.ShouldBindJSON(&professor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO professores (nome, email, cpf) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err := database.DB.QueryRow(query, professor.Nome, professor.Email, professor.CPF).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		professor.ID = id
		c.JSON(http.StatusOK, professor)
	})

	r.GET("/professores", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT id, nome, email, cpf FROM professores")
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

	// Turmas
	r.POST("/turmas", func(c *gin.Context) {
		var turma models.Turma
		if err := c.ShouldBindJSON(&turma); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se o professor existe
		var professorID int
		err := database.DB.QueryRow("SELECT id FROM professores WHERE id = $1", turma.ProfessorID).Scan(&professorID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Professor inválido"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar professor"})
			return
		}

		// Inserir a turma no banco de dados
		query := `INSERT INTO turmas (nome, ano, professor_id) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = database.DB.QueryRow(query, turma.Nome, turma.Ano, turma.ProfessorID).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		turma.ID = id
		c.JSON(http.StatusOK, turma)
	})

	r.GET("/turmas", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT id, nome, ano, professor_id FROM turmas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var turmas []models.Turma
		for rows.Next() {
			var turma models.Turma
			if err := rows.Scan(&turma.ID, &turma.Nome, &turma.Ano, &turma.ProfessorID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			turmas = append(turmas, turma)
		}

		c.JSON(http.StatusOK, turmas)
	})

	// Alunos
	// Alunos
	r.POST("/alunos", func(c *gin.Context) {
		var aluno models.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a turma existe
		var turmaID int
		err := database.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", aluno.TurmaID).Scan(&turmaID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turma"})
			return
		}

		// Inserir aluno no banco de dados
		query := `INSERT INTO alunos (nome, matricula, turma_id) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = database.DB.QueryRow(query, aluno.Nome, aluno.Matricula, aluno.TurmaID).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		aluno.ID = id
		c.JSON(http.StatusOK, aluno)
	})

	r.GET("/alunos", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT id, nome, matricula, turma_id FROM alunos")
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

	// Atividades
	r.POST("/atividades", func(c *gin.Context) {
		var atividade models.Atividade
		if err := c.ShouldBindJSON(&atividade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a soma dos valores das atividades da turma não ultrapassa 100 pontos
		var total float64
		query := `SELECT SUM(valor) FROM atividades WHERE turma_id = $1`
		err := database.DB.QueryRow(query, atividade.TurmaID).Scan(&total)
		if err != nil && err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		total += atividade.Valor
		if total > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O valor total das atividades da turma não pode ultrapassar 100 pontos"})
			return
		}

		query = `INSERT INTO atividades (turma_id, valor, data) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = database.DB.QueryRow(query, atividade.TurmaID, atividade.Valor, atividade.Data).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		atividade.ID = id
		c.JSON(http.StatusOK, atividade)
	})

	r.GET("/atividades", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT id, turma_id, valor, data FROM atividades")
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

	// Notas
	r.POST("/notas", func(c *gin.Context) {
		var nota models.Nota
		if err := c.ShouldBindJSON(&nota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a nota não ultrapassa o valor máximo permitido da atividade
		var valorAtividade float64
		query := `SELECT valor FROM atividades WHERE id = $1`
		err := database.DB.QueryRow(query, nota.AtividadeID).Scan(&valorAtividade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Atividade não encontrada"})
			return
		}

		if nota.Valor > valorAtividade {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A nota não pode ultrapassar o valor máximo da atividade"})
			return
		}

		query = `INSERT INTO notas (aluno_id, atividade_id, nota) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = database.DB.QueryRow(query, nota.AlunoID, nota.AtividadeID, nota.Valor).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		nota.ID = id
		c.JSON(http.StatusOK, nota)
	})

	r.GET("/notas", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT id, aluno_id, atividade_id, nota FROM notas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var notas []models.Nota
		for rows.Next() {
			var nota models.Nota
			if err := rows.Scan(&nota.ID, &nota.AlunoID, &nota.AtividadeID, &nota.Valor); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			notas = append(notas, nota)
		}

		c.JSON(http.StatusOK, notas)
	})
}
