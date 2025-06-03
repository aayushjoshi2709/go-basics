package main

import (
	"aayushjoshi2709/rest-apis/db"
	"aayushjoshi2709/rest-apis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
