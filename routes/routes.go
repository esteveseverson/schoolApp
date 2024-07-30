package routes

import (
	"context"
	"net/http"
	"schoolApp/database"
	"schoolApp/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupRoutes(r *gin.Engine) {
	// Professores
	r.POST("/professores", func(c *gin.Context) {
		var professor models.Professor
		if err := c.ShouldBindJSON(&professor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := database.DB.Collection("professores")
		counterCollection := database.DB.Collection("counters")

		id, err := database.GetNextSequence(counterCollection, "professores")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		professor.ID = id

		result, err := collection.InsertOne(context.Background(), professor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	// Turmas
	r.POST("/turmas", func(c *gin.Context) {
		var turma models.Turma
		if err := c.ShouldBindJSON(&turma); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := database.DB.Collection("turmas")
		counterCollection := database.DB.Collection("counters")

		id, err := database.GetNextSequence(counterCollection, "turmas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		turma.ID = id

		result, err := collection.InsertOne(context.Background(), turma)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	// Alunos
	r.POST("/alunos", func(c *gin.Context) {
		var aluno models.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := database.DB.Collection("alunos")
		counterCollection := database.DB.Collection("counters")

		id, err := database.GetNextSequence(counterCollection, "alunos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		aluno.ID = id

		result, err := collection.InsertOne(context.Background(), aluno)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	// Atividades
	r.POST("/atividades", func(c *gin.Context) {
		var atividade models.Atividade
		if err := c.ShouldBindJSON(&atividade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := database.DB.Collection("atividades")
		counterCollection := database.DB.Collection("counters")

		id, err := database.GetNextSequence(counterCollection, "atividades")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		atividade.ID = id

		// Verificar se a soma dos valores das atividades da turma não ultrapassa 100 pontos
		filter := bson.M{"turma_id": atividade.TurmaID}
		var atividades []models.Atividade
		cursor, err := collection.Find(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err = cursor.All(context.Background(), &atividades); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		total := atividade.Valor
		for _, a := range atividades {
			total += a.Valor
		}

		if total > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "O valor total das atividades da turma não pode ultrapassar 100 pontos"})
			return
		}

		result, err := collection.InsertOne(context.Background(), atividade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	// Notas
	r.POST("/notas", func(c *gin.Context) {
		var nota models.Nota
		if err := c.ShouldBindJSON(&nota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		atividadesCollection := database.DB.Collection("atividades")
		ctx := context.Background()

		// Verificar se a nota não ultrapassa o valor máximo permitido da atividade
		var atividade models.Atividade
		err := atividadesCollection.FindOne(ctx, bson.M{"id": nota.AtividadeID}).Decode(&atividade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Atividade não encontrada"})
			return
		}

		if nota.Nota > atividade.Valor {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A nota não pode ultrapassar o valor máximo da atividade"})
			return
		}

		collection := database.DB.Collection("notas")
		counterCollection := database.DB.Collection("counters")

		id, err := database.GetNextSequence(counterCollection, "notas")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		nota.ID = id

		result, err := collection.InsertOne(ctx, nota)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
