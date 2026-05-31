package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/clothing/image", postClothingImage)

	router.Run("localhost:8080")
}

func postClothingImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	destinationDir, err := filepath.Abs("../uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	destination := filepath.Join(destinationDir, filepath.Base(file.Filename))
	err = c.SaveUploadedFile(file, destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, "")
}
