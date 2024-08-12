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
			// Código existente para POST...
		})

		professores.PUT("/:id", func(c *gin.Context) {
			// Código existente para PUT...
		})

		professores.DELETE("/:id", func(c *gin.Context) {
			// Código existente para DELETE...
		})

		professores.GET("", func(c *gin.Context) {
			var professores []models.Professor
			collection := database.DB.Collection("professores")

			cursor, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch professors", err)
				return
			}

			if err := cursor.All(context.Background(), &professores); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to decode professors", err)
				return
			}

			c.JSON(http.StatusOK, professores)
		})
	}

	// Turmas
	turmas := r.Group("/turmas")
	{
		turmas.POST("", func(c *gin.Context) {
			// Código existente para POST...
		})

		turmas.PUT("/:id", func(c *gin.Context) {
			// Código existente para PUT...
		})

		turmas.DELETE("/:id", func(c *gin.Context) {
			// Código existente para DELETE...
		})

		turmas.GET("", func(c *gin.Context) {
			var turmas []models.Turma
			collection := database.DB.Collection("turmas")

			cursor, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch turmas", err)
				return
			}

			if err := cursor.All(context.Background(), &turmas); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to decode turmas", err)
				return
			}

			c.JSON(http.StatusOK, turmas)
		})
	}

	// Alunos
	alunos := r.Group("/alunos")
	{
		alunos.POST("", func(c *gin.Context) {
			// Código existente para POST...
		})

		alunos.PUT("/:id", func(c *gin.Context) {
			// Código existente para PUT...
		})

		alunos.DELETE("/:id", func(c *gin.Context) {
			// Código existente para DELETE...
		})

		alunos.GET("", func(c *gin.Context) {
			var alunos []models.Aluno
			collection := database.DB.Collection("alunos")

			cursor, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch alunos", err)
				return
			}

			if err := cursor.All(context.Background(), &alunos); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to decode alunos", err)
				return
			}

			c.JSON(http.StatusOK, alunos)
		})
	}

	// Atividades
	atividades := r.Group("/atividades")
	{
		atividades.POST("", func(c *gin.Context) {
			// Código existente para POST...
		})

		atividades.PUT("/:id", func(c *gin.Context) {
			// Código existente para PUT...
		})

		atividades.DELETE("/:id", func(c *gin.Context) {
			// Código existente para DELETE...
		})

		atividades.GET("", func(c *gin.Context) {
			var atividades []models.Atividade
			collection := database.DB.Collection("atividades")

			cursor, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch atividades", err)
				return
			}

			if err := cursor.All(context.Background(), &atividades); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to decode atividades", err)
				return
			}

			c.JSON(http.StatusOK, atividades)
		})
	}

	// Notas
	notas := r.Group("/notas")
	{
		notas.POST("", func(c *gin.Context) {
			// Código existente para POST...
		})

		notas.PUT("/:id", func(c *gin.Context) {
			// Código existente para PUT...
		})

		notas.DELETE("/:id", func(c *gin.Context) {
			// Código existente para DELETE...
		})

		notas.GET("", func(c *gin.Context) {
			var notas []models.Nota
			collection := database.DB.Collection("notas")

			cursor, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to fetch notas", err)
				return
			}

			if err := cursor.All(context.Background(), &notas); err != nil {
				utils.HandleError(c, http.StatusInternalServerError, "Failed to decode notas", err)
				return
			}

			c.JSON(http.StatusOK, notas)
		})
	}
}
