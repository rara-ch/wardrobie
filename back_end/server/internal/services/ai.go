package ai

import (
	"context"
	"os"

	"google.golang.org/genai"
)

type AiService struct {
	client *genai.Client
}

func NewService(client *genai.Client) *AiService {
	return &AiService{client: client}
}

func (s *AiService) AnalyseImage(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	parts := []*genai.Part{
		genai.NewPartFromBytes(bytes, "image/jpg"),
		genai.NewPartFromText("Describe this image"),
	}

	result, err := s.client.Models.GenerateContent(
		context.Background(),
		"gemini-3.5-flash",
		[]*genai.Content{genai.NewContentFromParts(parts, genai.RoleUser)},
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
