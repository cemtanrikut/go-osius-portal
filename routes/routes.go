package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/company", handlers.CreateCompany)
	r.GET("/companies", handlers.GetCompanies)
}
