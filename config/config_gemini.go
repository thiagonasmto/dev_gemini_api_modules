package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func ConfigGemini() (ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, debugResponse bool, erro error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY não está definido. Defina a chave de API como variável de ambiente.")
	}

	ctx = context.Background()
	client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	model = "gemini-2.5-flash"
	debugResponse = false

	budget := int32(-1) // Define o valor para a configuração do Pensamento do Gemini (0) Desabilita, (0 a 24576 para o modelo flash) Habilita e (-1) Dinâmico.
	config = &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget:  &budget,       // Habilita, Desabilita ou define como Dinâmico o Pensamento do Gemini
			IncludeThoughts: debugResponse, // Habilita ou Desabilita o Resumo de Ideias. Resumo de Ideias é a formulação da resposta da LLM.
		},
	}

	return ctx, client, model, config, debugResponse, nil
}
