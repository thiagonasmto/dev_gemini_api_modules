package controllers

import (
	"context"
	"geminai_with_go/models"
	utilexamples "geminai_with_go/services/util_examples"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
)

// controllers/simple_chat.go
func SimpleLocalChat(c *gin.Context, ctx_ollama context.Context, client_ollama *api.Client) {
	var request models.OllamaResquestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Stream para o cliente
	c.Stream(func(w io.Writer) bool {
		err := utilexamples.LocalSimpleRequest(ctx_ollama, client_ollama, request.Prompt, request.Model, func(chunk string) {
			// envia chunk em tempo real para o cliente
			c.SSEvent("message", chunk)
		})

		if err != nil {
			c.SSEvent("error", err.Error())
			return false
		}
		return false // false indica que acabou o streaming
	})
}
