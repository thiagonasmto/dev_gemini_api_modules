package utilfunctions

import (
	"context"

	"github.com/ollama/ollama/api"
)

func CreateEmbedding(ctx context.Context, client *api.Client, model_embedding, text string) ([]float32, error) {
	resp, err := client.Embed(ctx, &api.EmbedRequest{
		Model: model_embedding, // ex: "nomic-embed-text"
		Input: []string{text},
	})
	if err != nil {
		return nil, err
	}
	return resp.Embeddings[0], nil
}
