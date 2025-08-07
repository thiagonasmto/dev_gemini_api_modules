package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func ConfigGemini() (ctx context.Context, client *genai.Client, model string, erro error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY não está definido. Defina a chave de API como variável de ambiente.")
	}

	ctx = context.Background()
	client, err = genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		log.Fatal(err)
	}

	model = "gemini-2.5-flash"

	return ctx, client, model, nil
}
