package geminiresources

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

func UderstandingLocalImages(ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, prompt string, path_image string) (*genai.GenerateContentResponse, error) {
	bytes, _ := os.ReadFile(path_image)

	parts := []*genai.Part{
		genai.NewPartFromBytes(bytes, "image/jpeg"),
		genai.NewPartFromText(prompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		model,
		contents,
		config,
	)

	fmt.Println(result.Text())
	return result, nil
}
