package controllers

import (
	"context"
	"geminai_with_go/models"
	utilexamples "geminai_with_go/services/util_examples"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func ChatWithHistory(c *gin.Context, ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig) {
	var input models.ChatDTO

	// Bind dos campos JSON + form-data
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input json"})
		return
	}

	var history []*genai.Content

	for _, msg := range input.History {
		for _, p := range msg.Parts {
			if p.Type == "text" {
				history = append(history, genai.NewContentFromText(p.Data, genai.Role(msg.Role)))
			}
		}
	}

	// Chama Gemini
	response_llm, err := utilexamples.ChatGemini(ctx, client, history, model, config, input.Prompt)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get response from Gemini"})
		return
	}

	c.JSON(200, gin.H{
		"response": response_llm,
	})
}
