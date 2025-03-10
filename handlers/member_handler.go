package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"main.go/config"
// 	"main.go/models"
// )

// // CreateMember godoc
// // @Summary Creates a new member
// // @Description Adds a new member to the database
// // @Tags members
// // @Accept  json
// // @Produce  json
// // @Param   member  body     models.Member  true  "Member data"
// // @Success 201      {object} models.Member
// // @Failure 400      {object} map[string]string
// // @Router  /members [post]
// func CreateMember(c *gin.Context) {
// 	var member models.Member
// 	if err := c.ShouldBindJSON(&member); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Create(&member)
// 	c.JSON(http.StatusCreated, member)
// }

// // GetMembers godoc
// // @Summary Retrieves all members
// // @Description Gets a list of all members
// // @Tags members
// // @Produce  json
// // @Success 200  {array}  models.Member
// // @Failure 500  {object} map[string]string
// // @Router  /members [get]
// func GetMembers(c *gin.Context) {
// 	var members []models.Member
// 	config.DB.Find(&members)
// 	c.JSON(http.StatusOK, members)
// }

// // GetMemberByID godoc
// // @Summary Retrieves a member by ID
// // @Description Gets a specific member using its ID
// // @Tags members
// // @Produce  json
// // @Param   id   path     int  true  "Member ID"
// // @Success 200  {object} models.Member
// // @Failure 404  {object} map[string]string
// // @Router  /members/{id} [get]
// func GetMemberByID(c *gin.Context) {
// 	var member models.Member
// 	id := c.Param("id")

// 	if err := config.DB.First(&member, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, member)
// }

// // UpdateMember godoc
// // @Summary Updates a member
// // @Description Updates an existing member by ID
// // @Tags members
// // @Accept  json
// // @Produce  json
// // @Param   id        path     int              true  "Member ID"
// // @Param   member  body     models.Member  true  "Updated member data"
// // @Success 200      {object} models.Member
// // @Failure 400      {object} map[string]string
// // @Failure 404      {object} map[string]string
// // @Router  /members/{id} [put]
// func UpdateMember(c *gin.Context) {
// 	var member models.Member
// 	id := c.Param("id")

// 	if err := config.DB.First(&member, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&member); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	config.DB.Save(&member)
// 	c.JSON(http.StatusOK, member)
// }

// // DeleteMember godoc
// // @Summary Deletes a member
// // @Description Deletes a specific member using its ID
// // @Tags members
// // @Param   id   path     int  true  "Member ID"
// // @Success 200  {object} map[string]string
// // @Failure 404  {object} map[string]string
// // @Router  /members/{id} [delete]
// func DeleteMember(c *gin.Context) {
// 	var member models.Member
// 	id := c.Param("id")

// 	if err := config.DB.First(&member, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
// 		return
// 	}

// 	config.DB.Delete(&member)
// 	c.JSON(http.StatusOK, gin.H{"message": "Member deleted succesfully!"})
// }
