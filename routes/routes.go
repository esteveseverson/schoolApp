package routes

import (
	"context"
	"net/http"
	"schoolApp/database"
	"schoolApp/models"
	"schoolApp/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupRoutes(r *gin.Engine) {
	// Professores
	professores := r.Group("/professores")
	{
		professores.POST("", func(c *gin.Context) {
			var professor models.Professor
			if err := c.ShouldBindJSON(&professor); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			collection := database.DB.Collection("professores")
			counterCollection := database.DB.Collection("counters")

			id, err := database.GetNextSequence(counterCollection, "professores")
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha na sequência de IDs", err)
				return
			}

			professor.ID = id

			_, err = collection.InsertOne(context.Background(), professor)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao inserir porofessor", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   professor.ID,
				"nome": professor.Nome,
			})
		})

		professores.PUT("/:id", func(c *gin.Context) {
			var professor models.Professor
			if err := c.ShouldBindJSON(&professor); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			id := c.Param("id")
			collection := database.DB.Collection("professores")

			filter := bson.M{"id": id}
			update := bson.M{"$set": professor}

			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao atualizar professor", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   professor.ID,
				"nome": professor.Nome,
			})
		})

		professores.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			collection := database.DB.Collection("professores")

			filter := bson.M{"id": id}
			_, err := collection.DeleteOne(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao deletar professor", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Professor deletado com sucesso",
			})
		})
	}

	// Turmas
	turmas := r.Group("/turmas")
	{
		turmas.POST("", func(c *gin.Context) {
			var turma models.Turma
			if err := c.ShouldBindJSON(&turma); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			collection := database.DB.Collection("turmas")
			counterCollection := database.DB.Collection("counters")

			id, err := database.GetNextSequence(counterCollection, "turmas")
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha na sequência de IDs", err)
				return
			}

			turma.ID = id

			_, err = collection.InsertOne(context.Background(), turma)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao inserir turma", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   turma.ID,
				"nome": turma.Nome,
			})
		})

		turmas.PUT("/:id", func(c *gin.Context) {
			var turma models.Turma
			if err := c.ShouldBindJSON(&turma); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			id := c.Param("id")
			collection := database.DB.Collection("turmas")

			filter := bson.M{"id": id}
			update := bson.M{"$set": turma}

			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao atualizar turma", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   turma.ID,
				"nome": turma.Nome,
			})
		})

		turmas.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			collection := database.DB.Collection("turmas")

			filter := bson.M{"id": id}
			_, err := collection.DeleteOne(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao deletar turma", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Turma deletada com sucesso",
			})
		})
	}

	// Alunos
	alunos := r.Group("/alunos")
	{
		alunos.POST("", func(c *gin.Context) {
			var aluno models.Aluno
			if err := c.ShouldBindJSON(&aluno); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			collection := database.DB.Collection("alunos")
			counterCollection := database.DB.Collection("counters")

			id, err := database.GetNextSequence(counterCollection, "alunos")
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha na sequência de IDs", err)
				return
			}

			aluno.ID = id

			_, err = collection.InsertOne(context.Background(), aluno)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao inserir aluno", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   aluno.ID,
				"nome": aluno.Nome,
			})
		})

		alunos.PUT("/:id", func(c *gin.Context) {
			var aluno models.Aluno
			if err := c.ShouldBindJSON(&aluno); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			id := c.Param("id")
			collection := database.DB.Collection("alunos")

			filter := bson.M{"id": id}
			update := bson.M{"$set": aluno}

			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao atualizar aluno", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":   aluno.ID,
				"nome": aluno.Nome,
			})
		})

		alunos.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			collection := database.DB.Collection("alunos")

			filter := bson.M{"id": id}
			_, err := collection.DeleteOne(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao deletar aluno", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Aluno deletado com sucesso",
			})
		})
	}

	// Atividades
	atividades := r.Group("/atividades")
	{
		atividades.POST("", func(c *gin.Context) {
			var atividade models.Atividade
			if err := c.ShouldBindJSON(&atividade); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			collection := database.DB.Collection("atividades")
			counterCollection := database.DB.Collection("counters")

			id, err := database.GetNextSequence(counterCollection, "atividades")
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha na sequencia de IDs", err)
				return
			}

			atividade.ID = id

			// Verificar se a soma dos valores das atividades da turma não ultrapassa 100 pontos
			filter := bson.M{"turma_id": atividade.TurmaID}
			var atividades []models.Atividade
			cursor, err := collection.Find(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao encontrar atividade", err)
				return
			}
			if err = cursor.All(context.Background(), &atividades); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao carregar atividade", err)
				return
			}

			total := atividade.Valor
			for _, a := range atividades {
				total += a.Valor
			}

			if total > 100 {
				utils.HandleError(c, http.StatusBadRequest, "O valor total das atividades da turma não pode ultrapassar 100 pontos", nil)
				return
			}

			_, err = collection.InsertOne(context.Background(), atividade)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao inserir atividade", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id": atividade.ID,
			})
		})

		atividades.PUT("/:id", func(c *gin.Context) {
			var atividade models.Atividade
			if err := c.ShouldBindJSON(&atividade); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			id := c.Param("id")
			collection := database.DB.Collection("atividades")

			filter := bson.M{"id": id}
			update := bson.M{"$set": atividade}

			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao atualizar atividade", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id": atividade.ID,
			})
		})

		atividades.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			collection := database.DB.Collection("atividades")

			filter := bson.M{"id": id}
			_, err := collection.DeleteOne(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao deletar atividade", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Atividade deletada com sucesso",
			})
		})
	}

	// Notas
	notas := r.Group("/notas")
	{
		notas.POST("", func(c *gin.Context) {
			var nota models.Nota
			if err := c.ShouldBindJSON(&nota); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			collection := database.DB.Collection("notas")
			counterCollection := database.DB.Collection("counters")

			id, err := database.GetNextSequence(counterCollection, "notas")
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha na sequencia de IDs", err)
				return
			}

			nota.ID = id

			// Verificar se a nota não ultrapassa o valor da atividade
			atividadeCollection := database.DB.Collection("atividades")
			var atividade models.Atividade
			err = atividadeCollection.FindOne(context.Background(), bson.M{"id": nota.AtividadeID}).Decode(&atividade)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao encontrar atividade", err)
				return
			}

			if nota.Nota > atividade.Valor {
				utils.HandleError(c, http.StatusBadRequest, "Nota não pode ser maior que o valor da atividade", nil)
				return
			}

			_, err = collection.InsertOne(context.Background(), nota)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao inserir nota", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id": nota.ID,
			})
		})

		notas.PUT("/:id", func(c *gin.Context) {
			var nota models.Nota
			if err := c.ShouldBindJSON(&nota); err != nil {
				utils.HandleError(c, http.StatusBadRequest, "JSON inválido", err)
				return
			}

			id := c.Param("id")
			collection := database.DB.Collection("notas")

			filter := bson.M{"id": id}
			update := bson.M{"$set": nota}

			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao atualizar nota", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id": nota.ID,
			})
		})

		notas.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			collection := database.DB.Collection("notas")

			filter := bson.M{"id": id}
			_, err := collection.DeleteOne(context.Background(), filter)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Falha ao deletar nota", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Nota deletatada com sucesso",
			})
		})
	}
}
