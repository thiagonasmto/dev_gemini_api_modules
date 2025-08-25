package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"geminai_with_go/config"
	"geminai_with_go/models"
	utilfunctions "geminai_with_go/services/util_functions"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
	"github.com/pgvector/pgvector-go"
)

func PdfExtractor(c *gin.Context, ctx_ollama context.Context, client *api.Client, model_embedding string) error {
	// var Document models.Document
	response, err := utilfunctions.ExtractTextFromPDF("./docs/pdf_examples/webinar-ibef-sobena-apresentacao.pdf")
	if err != nil {
		fmt.Print("erro ao extrair texto do arquivo pdf: ", err)
	}
	fmt.Println(response)
	chunks := utilfunctions.ChunkText(response, "", 500)

	for i, chunk := range chunks {
		embedding, err := utilfunctions.CreateEmbedding(ctx_ollama, client, model_embedding, chunk.Content)
		if err != nil {
			fmt.Print("erro ao gerar embedding: ", err)
		}
		fmt.Printf("Embedding %d: %v\n", i, embedding)
		metadata, _ := json.Marshal(chunk.Metadata)

		embeddingVec := pgvector.NewVector(embedding)
		doc := models.Document{
			ID:        chunk.ID,
			Content:   chunk.Content,
			Metadata:  metadata,
			Embedding: embeddingVec,
		}

		if err := config.DB.Create(&doc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro ao salvar embedding": err.Error()})
			return err
		}
	}

	return nil
}
