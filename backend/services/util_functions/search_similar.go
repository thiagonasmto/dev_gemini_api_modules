package utilfunctions

import (
	"context"
	"fmt"
	"geminai_with_go/models"

	"github.com/ollama/ollama/api"
)

func SearchSimilar(ctx_ollama context.Context, client *api.Client, model_embedding string, k int, prompt string) ([]models.Document, error) {
	embedding, err := CreateEmbedding(ctx_ollama, client, model_embedding, prompt)
	if err != nil {
		fmt.Println("erro ao criar embedding do prompt do usuário: ", err)
		return nil, err
	}
	return embedding, nil
}
