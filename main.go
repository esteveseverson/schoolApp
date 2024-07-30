package main

import (
	"schoolApp/database"
	"schoolApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
