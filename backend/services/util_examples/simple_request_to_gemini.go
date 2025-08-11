package utilexamples

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func SimpleRequest(client *genai.Client, ctx context.Context, model string, prompt string, config *genai.GenerateContentConfig) (*genai.GenerateContentResponse, error) {
	contents := genai.Text(prompt)
	response, err := client.Models.GenerateContent(
		ctx,
		model,
		contents,
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	println(response.Text())
	return response, nil
}

func JSONSimpleRequest(response *genai.GenerateContentResponse) {
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))
}
