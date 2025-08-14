package geminiresources

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

func UnderstandingInLineDocuments(ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, prompt string, pdfBytes []byte) (*genai.GenerateContentResponse, error) {
	if pdfBytes == nil {
		return nil, fmt.Errorf("erro: nenhum arquivo enviado")
	}

	parts := []*genai.Part{
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "application/pdf",
				Data:     pdfBytes,
			},
		},
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
