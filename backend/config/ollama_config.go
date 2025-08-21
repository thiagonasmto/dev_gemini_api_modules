package config

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func OllamaConfig() (ctx context.Context, client *api.Client, erro error) {
	ctx = context.Background()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		fmt.Println("erro ao configurar cliente ollama: ", err)
		return nil, nil, err
	}

	return ctx, client, nil
}
