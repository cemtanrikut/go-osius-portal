package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"main.go/config"
// 	"main.go/models"
// )

// // CreateBuilding godoc
// // @Summary Creates a new building
// // @Description Adds a new building to the database
// // @Tags buildings
// // @Accept  json
// // @Produce  json
// // @Param   building  body     models.Building  true  "Building data"
// // @Success 201      {object} models.Building
// // @Failure 400      {object} map[string]string
// // @Router  /buildings [post]
// func CreateBuilding(c *gin.Context) {
// 	var building models.Building
// 	if err := c.ShouldBindJSON(&building); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Create(&building)
// 	c.JSON(http.StatusCreated, building)
// }

// // GetBuildings godoc
// // @Summary Retrieves all buildings
// // @Description Gets a list of all buildings
// // @Tags buildings
// // @Produce  json
// // @Success 200  {array}  models.Building
// // @Failure 500  {object} map[string]string
// // @Router  /buildings [get]
// func GetBuildings(c *gin.Context) {
// 	var buildings []models.Building
// 	config.DB.Find(&buildings)
// 	c.JSON(http.StatusOK, buildings)
// }

// // GetBuildingByID godoc
// // @Summary Retrieves a building by ID
// // @Description Gets a specific building using its ID
// // @Tags buildings
// // @Produce  json
// // @Param   id   path     int  true  "Building ID"
// // @Success 200  {object} models.Building
// // @Failure 404  {object} map[string]string
// // @Router  /buildings/{id} [get]
// func GetBuildingByID(c *gin.Context) {
// 	var building models.Building
// 	id := c.Param("id")

// 	if err := config.DB.First(&building, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, building)
// }

// // UpdateBuilding godoc
// // @Summary Updates a building
// // @Description Updates an existing building by ID
// // @Tags buildings
// // @Accept  json
// // @Produce  json
// // @Param   id        path     int              true  "Building ID"
// // @Param   building  body     models.Building  true  "Updated building data"
// // @Success 200      {object} models.Building
// // @Failure 400      {object} map[string]string
// // @Failure 404      {object} map[string]string
// // @Router  /buildings/{id} [put]
// func UpdateBuilding(c *gin.Context) {
// 	var building models.Building
// 	id := c.Param("id")

// 	if err := config.DB.First(&building, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&building); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Save(&building)
// 	c.JSON(http.StatusOK, building)
// }

// // DeleteBuilding godoc
// // @Summary Deletes a building
// // @Description Deletes a specific building using its ID
// // @Tags buildings
// // @Param   id   path     int  true  "Building ID"
// // @Success 200  {object} map[string]string
// // @Failure 404  {object} map[string]string
// // @Router  /buildings/{id} [delete]
// func DeleteBuilding(c *gin.Context) {
// 	var building models.Building
// 	id := c.Param("id")

// 	if err := config.DB.First(&building, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
// 		return
// 	}

// 	config.DB.Delete(&building)
// 	c.JSON(http.StatusOK, gin.H{"message": "Building deleted succesfully!"})
// }
