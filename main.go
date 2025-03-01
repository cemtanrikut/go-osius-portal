package main

import (
	_ "main.go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/config"
	"main.go/routes"
)

// @title ERP Panel API
// @version 1.0
// @description REST API documentation for ERP panel.
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r)

	r.Run(":8080")
}
