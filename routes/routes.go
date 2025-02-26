package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func SetupRoutes(r *gin.Engine) {
	// Company
	r.POST("/companies", handlers.CreateCompany)
	r.GET("/companies", handlers.GetCompanies)
	r.GET("/companies/:id", handlers.GetCompanyByID)
	r.PUT("/companies/:id", handlers.UpdateCompany)
	r.DELETE("/companies/:id", handlers.DeleteCompany)

	// Building
	r.POST("/buildings", handlers.CreateBuilding)
	r.GET("/buildings", handlers.GetBuildings)
	r.GET("/buildings/:id", handlers.GetBuildingByID)
	r.PUT("/buildings/:id", handlers.UpdateBuilding)
	r.DELETE("/buildings/:id", handlers.DeleteBuilding)

	// Room
	r.POST("/rooms", handlers.CreateRoom)
	r.GET("/rooms", handlers.GetRooms)
	r.GET("/rooms/:id", handlers.GetRoomByID)
	r.PUT("/rooms/:id", handlers.UpdateRoom)
	r.DELETE("/rooms/:id", handlers.DeleteRoom)

	// Member
	r.POST("/members", handlers.CreateMember)
	r.GET("/members", handlers.GetMembers)
	r.GET("/members/:id", handlers.GetMemberByID)
	r.PUT("/members/:id", handlers.UpdateMember)
	r.DELETE("/members/:id", handlers.DeleteMember)
}
