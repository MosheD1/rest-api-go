package main

import (
	"courseProject/db"
	"courseProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
