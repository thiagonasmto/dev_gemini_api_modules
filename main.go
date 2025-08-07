package main

import (
	"fmt"
	"geminai_with_go/config"
	utilexamples "geminai_with_go/util_examples"

	"google.golang.org/genai"
)

func main() {
	ctx, client, model, config, debugResponse, _ := config.ConfigGemini()
	prompt := "Explain how AI works in few words"
	content := genai.Text(prompt)

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
