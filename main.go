package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY não está definido. Defina a chave de API como variável de ambiente.")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		log.Fatal(err)
	}

	model := "gemini-2.5-flash"
	prompt := "Explain how AI works in a few words"
	content := genai.Text(prompt)

	budget := int32(-1) // Define o valor para a configuração do Pensamento do Gemini (0) Desabilita, (0 a 24576 para o modelo flash) Habilita e (-1) Dinâmico.
	config := &genai.GenerateContentConfig{
		// ResponseMIMEType: "application/json",
		// ResponseSchema: &genai.Schema{
		// 	Type: genai.TypeArray,
		// 	Items: &genai.Schema{
		// 		Type: genai.TypeObject,
		// 		Properties: map[string]*genai.Schema{
		// 			"recipeName": {Type: genai.TypeString},
		// 			"ingredients": {
		// 				Type:  genai.TypeArray,
		// 				Items: &genai.Schema{Type: genai.TypeString},
		// 			},
		// 		},
		// 		PropertyOrdering: []string{"recipeName", "ingredients"},
		// 	},
		// },
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget:  &budget, // Habilita, Desabilita ou define como Dinâmico o Pensamento do Gemini
			IncludeThoughts: false,   // Habilita ou Desabilita o Resumo de Ideias. Resumo de Ideias é a formulação da resposta da LLM.
		},
	}

	response, err := client.Models.GenerateContent(
		ctx,
		model,
		content,
		// nil,
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	// for _, part := range response.Candidates[0].Content.Parts {
	// 	if part.Text != "" {
	// 		if part.Thought {
	// 			fmt.Println("Thoughts Summary:")
	// 			fmt.Println(part.Text)
	// 		} else {
	// 			fmt.Println("Answer:")
	// 			fmt.Println(part.Text)
	// 		}
	// 	}
	// }

	// ...
	// usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Thoughts tokens:", string(usageMetadata.thoughts_token_count))
	// fmt.Println("Output tokens:", string(usageMetadata.candidates_token_count))

	// response com streaming
	// response_stream := client.Models.GenerateContentStream(
	// 	ctx,
	// 	model,
	// 	content,
	// 	&genai.GenerateContentConfig{
	// 		ThinkingConfig: &genai.ThinkingConfig{
	// 			IncludeThoughts: true,
	// 		},
	// 	},
	// )

	// for chunk := range response_stream {
	// 	for _, part := range chunk.Candidates[0].Content.Parts {
	// 		if len(part.Text) == 0 {
	// 			continue
	// 		}

	// 		if part.Thought {
	// 			fmt.Printf("Thought: %s\n", part.Text)
	// 		} else {
	// 			fmt.Printf("Answer: %s\n", part.Text)
	// 		}
	// 	}
	// }

	// fmt.Print(response.SDKHTTPResponse)
	// fmt.Print(response.UsageMetadata)
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))
	// fmt.Print(response.Text())
}
