package main

import (
	_ "main.go/docs"
	"main.go/handlers"

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

	// r := gin.Default()

	// // Swagger UI
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := routes.SetupRouter()

	go handlers.BroadcastMessages() // Mesajları dinlemeye başla

	r.Run(":8080")
}
