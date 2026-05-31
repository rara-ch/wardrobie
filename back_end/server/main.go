package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type state struct {
	gemini *genai.Client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file: %s", err)
	}

	context := context.Background()
	gemini, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatalf("could not create an gemini client: %s", err)
	}

	state := &state{
		gemini: gemini,
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

	bytes, err := os.ReadFile(fmt.Sprintf("../uploads/%s", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	parts := []*genai.Part{
		genai.NewPartFromBytes(bytes, "image/jpg"),
		genai.NewPartFromText("Describe this image"),
	}

	result, err := state.gemini.Models.GenerateContent(
		context.Background(),
		"gemini-3.5-flash",
		[]*genai.Content{genai.NewContentFromParts(parts, genai.RoleUser)},
		nil,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"analysis": result.Text()})
}
