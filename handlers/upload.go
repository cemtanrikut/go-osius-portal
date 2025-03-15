package handlers

// UploadFile godoc
// @Summary Uploads a file
// @Description Handles file uploads for messages
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /upload [post]
// func UploadFile(c *gin.Context) {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
// 		return
// 	}

// 	uploadPath := "uploads/"
// 	filename := filepath.Join(uploadPath, file.Filename)

// 	if err := c.SaveUploadedFile(file, filename); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
// 		return
// 	}

// 	fileURL := fmt.Sprintf("/uploads/%s", file.Filename)
// 	c.JSON(http.StatusOK, gin.H{"file_url": fileURL})
// }
