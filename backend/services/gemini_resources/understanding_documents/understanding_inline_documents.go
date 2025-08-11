package geminiresources

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/genai"
)

func UnderstandingInLineDocuments(ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, prompt string, doc_url string) (*genai.GenerateContentResponse, error) {
	pdfResp, err := http.Get(doc_url)
	if err != nil {
		return nil, fmt.Errorf("erro: falha na requisição para captura do arquivo")
	}
	var pdfBytes []byte
	if pdfResp != nil && pdfResp.Body != nil {
		pdfBytes, _ = io.ReadAll(pdfResp.Body)
		pdfResp.Body.Close()
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
