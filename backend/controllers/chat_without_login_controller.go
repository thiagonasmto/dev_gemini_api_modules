package controllers

import (
	"context"
	"geminai_with_go/models"
	utilexamples "geminai_with_go/services/util_examples"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func SimpleRequestController(c *gin.Context, ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, debugResponse bool) {
	var input models.MessageDTO

	// Bind JSON da requisição para a struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	// Chama o serviço que faz a requisição ao Gemini
	response_llm, err := utilexamples.SimpleRequest(client, ctx, model, input.Content, config)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get response from Gemini"})
		return
	}

	if debugResponse {
		utilexamples.SummaryResponse(response_llm)
	}

	// Retorna a resposta para o cliente no JSON
	c.JSON(200, gin.H{
		"response": response_llm,
	})
}
