package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"main.go/config"
// 	"main.go/models"
// )

// // CreateRoom godoc
// // @Summary Creates a new room
// // @Description Adds a new room to the database
// // @Tags rooms
// // @Accept  json
// // @Produce  json
// // @Param   room  body     models.Room  true  "Room data"
// // @Success 201      {object} models.Room
// // @Failure 400      {object} map[string]string
// // @Router  /rooms [post]
// func CreateRoom(c *gin.Context) {
// 	var room models.Room
// 	if err := c.ShouldBindJSON(&room); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Create(&room)
// 	c.JSON(http.StatusCreated, room)
// }

// // GetRooms godoc
// // @Summary Retrieves all rooms
// // @Description Gets a list of all rooms
// // @Tags rooms
// // @Produce  json
// // @Success 200  {array}  models.Room
// // @Failure 500  {object} map[string]string
// // @Router  /rooms [get]
// func GetRooms(c *gin.Context) {
// 	var rooms []models.Room
// 	config.DB.Find(&rooms)
// 	c.JSON(http.StatusOK, rooms)
// }

// // GetRoomByID godoc
// // @Summary Retrieves a room by ID
// // @Description Gets a specific room using its ID
// // @Tags rooms
// // @Produce  json
// // @Param   id   path     int  true  "Room ID"
// // @Success 200  {object} models.Room
// // @Failure 404  {object} map[string]string
// // @Router  /rooms/{id} [get]
// func GetRoomByID(c *gin.Context) {
// 	var room models.Room
// 	id := c.Param("id")

// 	if err := config.DB.First(&room, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, room)
// }

// // UpdateRoom godoc
// // @Summary Updates a room
// // @Description Updates an existing room by ID
// // @Tags rooms
// // @Accept  json
// // @Produce  json
// // @Param   id        path     int              true  "Room ID"
// // @Param   room      body     models.Room  true  "Updated room data"
// // @Success 200      {object} models.Room
// // @Failure 400      {object} map[string]string
// // @Failure 404      {object} map[string]string
// // @Router  /rooms/{id} [put]
// func UpdateRoom(c *gin.Context) {
// 	var room models.Room
// 	id := c.Param("id")

// 	if err := config.DB.First(&room, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&room); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Save(&room)
// 	c.JSON(http.StatusOK, room)
// }

// // DeleteRoom godoc
// // @Summary Deletes a room
// // @Description Deletes a specific room using its ID
// // @Tags rooms
// // @Param   id   path     int  true  "Room ID"
// // @Success 200  {object} map[string]string
// // @Failure 404  {object} map[string]string
// // @Router  /rooms/{id} [delete]
// func DeleteRoom(c *gin.Context) {
// 	var room models.Room
// 	id := c.Param("id")

// 	if err := config.DB.First(&room, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
// 		return
// 	}

// 	config.DB.Delete(&room)
// 	c.JSON(http.StatusOK, gin.H{"message": "Room deleted succesfull!"})
// }
