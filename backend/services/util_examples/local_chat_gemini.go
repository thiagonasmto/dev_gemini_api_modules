package utilexamples

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

// services/util_examples/local_request.go
func LocalSimpleRequest(ctx_ollama context.Context, client_ollama *api.Client, prompt string, model string, onChunk func(string)) error {
	req := &api.GenerateRequest{
		Model:  model,
		Prompt: prompt,
	}

	respFunc := func(resp api.GenerateResponse) error {
		// printa no terminal
		fmt.Print(resp.Response)

		// também envia pro callback (controller pode mandar pro cliente)
		if onChunk != nil {
			onChunk(resp.Response)
		}
		return nil
	}

	if err := client_ollama.Generate(ctx_ollama, req, respFunc); err != nil {
		return err
	}

	fmt.Println()
	return nil
}
