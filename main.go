package main

import (
	"fmt"
	"geminai_with_go/config"
	utilexamples "geminai_with_go/util_examples"

	"google.golang.org/genai"
)

func main() {
	ctx, client, model, _ := config.ConfigGemini()
	prompt := "Explain how AI works in few words"
	content := genai.Text(prompt)
	debugResponse := true

	budget := int32(-1) // Define o valor para a configuração do Pensamento do Gemini (0) Desabilita, (0 a 24576 para o modelo flash) Habilita e (-1) Dinâmico.
	config := &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget:  &budget,       // Habilita, Desabilita ou define como Dinâmico o Pensamento do Gemini
			IncludeThoughts: debugResponse, // Habilita ou Desabilita o Resumo de Ideias. Resumo de Ideias é a formulação da resposta da LLM.
		},
	}

	println("=================== no streaming response ==================")
	// response sem streaming
	response, err := utilexamples.SimpleRequest(client, ctx, model, prompt, content, config)
	if err != nil {
		fmt.Println("erro na comunicação com o gemini: ", err)
	}

	// Summary da resposta
	if debugResponse {
		utilexamples.SummaryResponse(response)
	} else {
		fmt.Print(response.Text())
	}

	// JSON da resposta
	utilexamples.JSONSimpleRequest(response)

	println("==================== streaming response ====================")
	// response com streaming
	utilexamples.StreamingRequest(client, ctx, model, prompt, content, config)

	// geminiresources.UnderstandingDocuments()
}
