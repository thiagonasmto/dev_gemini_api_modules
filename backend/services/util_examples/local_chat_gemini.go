package utilexamples

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func LocalSimpleRequest() {
	ctx := context.Background()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	req := &api.GenerateRequest{
		Model:  "gemma3:1b",
		Prompt: "Escreva uma história sobre uma mochila mágica.",
	}

	var fullResponse string

	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		fullResponse += resp.Response // concatena cada parte da resposta
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(fullResponse)
}
