package geminiresources

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func UnderstandingDocuments() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY não está definido. Defina a chave de API como variável de ambiente.")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	pdfResp, _ := http.Get("https://discovery.ucl.ac.uk/id/eprint/10089234/1/343019_3_art_0_py4t4l_convrt.pdf")
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
		genai.NewPartFromText("Summarize this document"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		contents,
		nil,
	)

	fmt.Println(result.Text())
}
