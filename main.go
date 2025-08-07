package main

import (
	"fmt"
	"geminai_with_go/config"
	geminiresources "geminai_with_go/gemini_resources"
	utilexamples "geminai_with_go/util_examples"

	"google.golang.org/genai"
)

func main() {
	ctx, client, model, config, debugResponse, _ := config.ConfigGemini()
	prompt := "Explain how AI works in few words"
	contents := genai.Text(prompt)

	println("=================== no streaming response ==================")
	// response sem streaming
	// response, err := utilexamples.SimpleRequest(client, ctx, model, prompt, config)
	// if err != nil {
	// 	fmt.Println("erro na comunicação com o gemini: ", err)
	// }

	// // Summary da resposta
	// if debugResponse {
	// 	utilexamples.SummaryResponse(response)
	// } else {
	// 	fmt.Print(response.Text())
	// }

	// // JSON da resposta
	// utilexamples.JSONSimpleRequest(response)

	println("==================== streaming response ====================")
	// response com streaming
	// utilexamples.StreamingRequest(client, ctx, model, prompt, config)

	println("==================== understanding local documents ====================")
	// Entendendo documentos locais
	prompt = "Sobre o que fala esses documentos? Me explique em poucas palavras."
	pdfPaths := []string{
		"./docs/pdf_examples/webinar-ibef-sobena-apresentacao.pdf",
	}
	contents = geminiresources.UnderstandingLocalDocuments(pdfPaths, ctx, client, prompt)
	response, err := utilexamples.SimpleRequest(client, ctx, model, contents, config)
	if err != nil {
		fmt.Println("erro na comunicação com o gemini: ", err)
	}
	if debugResponse {
		utilexamples.SummaryResponse(response)
	} else {
		fmt.Print(response.Text())
	}
}
