package controllers

import (
	"context"
	"fmt"
	"geminai_with_go/models"
	geminiresources "geminai_with_go/services/gemini_resources/understanding_documents"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func ChatDocumentController(c *gin.Context, ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig) {
	var input models.MessageDTO
	var fileBytes []byte

	// Pega o arquivo enviado pelo usuário
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("requisição sem arquivo: ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": "arquivo não enviado"})
		// return
	} else {
		// Abre o arquivo
		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao abrir arquivo"})
			return
		}
		defer f.Close()

		// Lê bytes do arquivo
		fileBytes, err = io.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao ler arquivo"})
			return
		}
	}

	if input.Content == "" {
		input.Content = "Por favor, explique o conteúdo deste documento."
	}

	result, err := geminiresources.UnderstandingInLineDocuments(ctx, client, model, config, input.Content, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}
