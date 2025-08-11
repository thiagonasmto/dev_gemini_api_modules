package geminiresources

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/genai"
)

func UderstandingInLineImages(ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, prompt string, url_image string) (*genai.GenerateContentResponse, error) {
	// Download the image.
	imageResp, err := http.Get(url_image)
	if err != nil {
		return nil, fmt.Errorf("erro: falha ao tentar realizar download através do link")
	}

	imageBytes, err := io.ReadAll(imageResp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro: falha ao tentar converter imagem")
	}

	parts := []*genai.Part{
		genai.NewPartFromBytes(imageBytes, "image/jpeg"),
		genai.NewPartFromText(prompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		model,
		contents,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("erro: falha na geração de resposta do gemini")
	}

	fmt.Println(result.Text())
	return result, nil
}
