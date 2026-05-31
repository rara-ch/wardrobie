package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	ai "wardrobie/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type state struct {
	aiService *ai.AiService
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file: %s", err)
	}

	context := context.Background()
	geminiClient, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatalf("could not create an gemini client: %s", err)
	}

	state := &state{
		aiService: ai.NewService(geminiClient),
	}

	router := gin.Default()
	router.POST("/clothing/image", state.postClothingImage)

	router.Run("localhost:8080")
}

func (state *state) postClothingImage(c *gin.Context) {
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

	analysis, err := state.aiService.AnalyseImage(destination)

	c.JSON(http.StatusOK, gin.H{"analysis": analysis})
}
