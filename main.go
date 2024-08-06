package main

import (
	"log"
	"schoolApp/database"
	"schoolApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
