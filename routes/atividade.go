package routes

import (
	"database/sql"
	"net/http"
	"schoolApp/config"
	"schoolApp/middleware"
	"schoolApp/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AtividadeRoutes(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	// Rota para listar todas as atividades
	r.GET("/atividades", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, turma_id, valor, data_entrega FROM atividades")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var atividades []models.Atividade
		for rows.Next() {
			var (
				id          int
				turmaID     int
				valor       float64
				dataEntrega time.Time
			)
			if err := rows.Scan(&id, &turmaID, &valor, &dataEntrega); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			atividade := models.Atividade{
				ID:          id,
				TurmaID:     turmaID,
				Valor:       valor,
				DataEntrega: models.CustomDate{Time: dataEntrega},
			}
			atividades = append(atividades, atividade)
		}
		c.JSON(http.StatusOK, atividades)
	})

	// Rota para buscar uma atividade pelo ID
	r.GET("/atividades/:id", func(c *gin.Context) {
		id := c.Param("id")

		var atividade models.Atividade
		var dataEntrega time.Time

		query := "SELECT id, turma_id, valor, data_entrega FROM atividades WHERE id = $1"
		err := config.DB.QueryRow(query, id).Scan(&atividade.ID, &atividade.TurmaID, &atividade.Valor, &dataEntrega)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		atividade.DataEntrega = models.CustomDate{Time: dataEntrega}
		c.JSON(http.StatusOK, atividade)
	})

	// Rota para criar uma nova atividade
	r.POST("/atividades", func(c *gin.Context) {
		var atividade models.Atividade
		if err := c.ShouldBindJSON(&atividade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verificar se a turma existe
		var turmaID int
		err := config.DB.QueryRow("SELECT id FROM turmas WHERE id = $1", atividade.TurmaID).Scan(&turmaID)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Turma inválida"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turma"})
			return
		}

		// Verificar a soma das atividades já cadastradas
		var totalAtividades float64
		err = config.DB.QueryRow("SELECT COALESCE(SUM(valor), 0) FROM atividades WHERE turma_id = $1", atividade.TurmaID).Scan(&totalAtividades)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar a soma das atividades"})
			return
		}

		if totalAtividades+atividade.Valor > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A soma das atividades ultrapassa 100 pontos"})
			return
		}

		// Convertendo a data para o formato adequado (YYYY-MM-DD)
		dataEntregaFormatted := atividade.DataEntrega.Format("2006-01-02")

		// Inserir a nova atividade
		query := `INSERT INTO atividades (turma_id, valor, data_entrega) VALUES ($1, $2, $3) RETURNING id`
		var id int
		err = config.DB.QueryRow(query, atividade.TurmaID, atividade.Valor, dataEntregaFormatted).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		atividade.ID = id
		c.JSON(http.StatusOK, atividade)
	})

	// Rota para atualizar uma atividade existente
	r.PUT("/atividades/:id", func(c *gin.Context) {
		id := c.Param("id")
		var atividade models.Atividade
		if err := c.ShouldBindJSON(&atividade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convertendo a data para o formato adequado (YYYY-MM-DD)
		dataEntregaFormatted := atividade.DataEntrega.Format("2006-01-02")

		query := `UPDATE atividades SET turma_id=$1, valor=$2, data_entrega=$3 WHERE id=$4`
		_, err := config.DB.Exec(query, atividade.TurmaID, atividade.Valor, dataEntregaFormatted, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Atividade atualizada com sucesso"})
	})

	// Rota para deletar uma atividade existente
	r.DELETE("/atividades/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := `DELETE FROM atividades WHERE id=$1`
		_, err := config.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Atividade excluída com sucesso"})
	})
}
