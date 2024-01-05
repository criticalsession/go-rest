package main

import (
	"github.com/criticalsession/go-rest/db"
	"github.com/criticalsession/go-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
