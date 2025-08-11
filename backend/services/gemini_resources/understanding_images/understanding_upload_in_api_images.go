package geminiresources

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func UnderstandingUploadImage(ctx context.Context, client *genai.Client, config *genai.GenerateContentConfig, model string, prompt string, image_paths []string) (*genai.GenerateContentResponse, error) {
	var promptParts []*genai.Part

	for _, path := range image_paths {
		uploadedFile, err := client.Files.UploadFromPath(ctx, path, nil)
		if err != nil {
			log.Printf("Erro ao enviar arquivo %s: %v", path, err)
			continue
		}
		part := genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType)
		promptParts = append(promptParts, part)
	}

	promptParts = append(promptParts, genai.NewPartFromText(prompt))

	contents := []*genai.Content{
		genai.NewContentFromParts(promptParts, genai.RoleUser),
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
