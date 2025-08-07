package geminiresources

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func UnderstandingLocalDocuments(pdfPaths []string, ctx context.Context, client *genai.Client, prompt string) (contents []*genai.Content) {
	// Exaple pdfPaths
	// pdfPaths = []string{
	// 	"/path/to/file1.pdf",
	// 	"/path/to/file2.pdf",
	// 	"/path/to/file3.pdf",
	// }

	var promptParts []*genai.Part

	uploadConfig := &genai.UploadFileConfig{MIMEType: "application/pdf"}

	for _, path := range pdfPaths {
		uploadedFile, err := client.Files.UploadFromPath(ctx, path, uploadConfig)
		if err != nil {
			log.Printf("Erro ao enviar arquivo %s: %v", path, err)
			continue
		}
		part := genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType)
		promptParts = append(promptParts, part)
	}

	// Acrescenta a pergunta ao final do prompt
	promptParts = append(promptParts, genai.NewPartFromText(prompt))

	contents = []*genai.Content{
		genai.NewContentFromParts(promptParts, genai.RoleUser),
	}

	return contents
}
