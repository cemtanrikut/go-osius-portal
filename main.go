package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/routes"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
