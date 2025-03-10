package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 📌 **Dashboard Verilerini Getirme**
func GetDashboardData(c *gin.Context) {
	var todoCount, inProgressCount, doneCount int64
	config.DB.Model(&models.Ticket{}).Where("status = ?", "To Do").Count(&todoCount)
	config.DB.Model(&models.Ticket{}).Where("status = ?", "In Progress").Count(&inProgressCount)
	config.DB.Model(&models.Ticket{}).Where("status = ?", "Done").Count(&doneCount)

	// 📌 **Kullanıcı Bazlı Ticket Sayıları**
	var users []string
	config.DB.Model(&models.Ticket{}).Distinct("assigned_to").Pluck("assigned_to", &users)

	userTickets := []gin.H{}
	for _, user := range users {
		var todo, inProgress, done int64
		config.DB.Model(&models.Ticket{}).Where("assigned_to = ? AND status = ?", user, "To Do").Count(&todo)
		config.DB.Model(&models.Ticket{}).Where("assigned_to = ? AND status = ?", user, "In Progress").Count(&inProgress)
		config.DB.Model(&models.Ticket{}).Where("assigned_to = ? AND status = ?", user, "Done").Count(&done)

		userTickets = append(userTickets, gin.H{
			"user":       user,
			"todo":       todo,
			"inProgress": inProgress,
			"done":       done,
		})
	}

	// 📌 **JSON Yanıtı Gönder**
	c.JSON(http.StatusOK, gin.H{
		"totalTickets": todoCount + inProgressCount + doneCount,
		"ticketData": gin.H{
			"todo":       todoCount,
			"inProgress": inProgressCount,
			"done":       doneCount,
		},
		"userTickets": userTickets,
	})
}
