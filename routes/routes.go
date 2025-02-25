package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/companies", handlers.CreateCompany)
	r.GET("/companies", handlers.GetCompanies)
	r.GET("/companies/:id", handlers.GetCompanyByID)
	r.PUT("/companies/:id", handlers.UpdateCompany)
	r.DELETE("/companies/:id", handlers.DeleteCompany)
}
