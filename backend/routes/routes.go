package routes

import (
	"context"
	"geminai_with_go/controllers"

	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
	"google.golang.org/genai"
)

func ApiRoutes(router *gin.Engine, ctx_gemini context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, debugResponse bool, ctx_ollama context.Context, client_ollama *api.Client) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.POST("/", controllers.CreateClient)
		clientGroup.GET("/", controllers.GetClients)
		clientGroup.GET("/:id", controllers.GetClient)
		clientGroup.PUT("/:id", controllers.UpdateClient)
		clientGroup.DELETE("/:id", controllers.DeleteClient)
	}

	admGroup := router.Group("/adms")
	{
		admGroup.POST("/", controllers.CreateAdm)
		admGroup.GET("/", controllers.GetAdms)
		admGroup.GET("/:id", controllers.GetAdm)
		admGroup.PUT("/:id", controllers.UpdateAdm)
		admGroup.DELETE("/:id", controllers.DeleteAdm)
	}

	geminiGroup := router.Group("/gemini-service")
	{
		geminiGroup.POST("/", func(c *gin.Context) {
			controllers.SimpleRequestController(c, ctx_gemini, client, model, config, debugResponse)
		})
		geminiGroup.POST("/gemini-history", func(c *gin.Context) {
			controllers.ChatWithHistory(c, ctx_gemini, client, model, config)
		})
		geminiGroup.POST("/gemini-understanding-doc", func(c *gin.Context) {
			controllers.ChatDocumentController(c, ctx_gemini, client, model, config)
		})
	}

	ollamaGroup := router.Group("/ollama-service")
	{
		ollamaGroup.POST("/", func(ctx *gin.Context) {
			controllers.SimpleLocalChat(ctx, ctx_ollama, client_ollama)
		})
		ollamaGroup.POST("/extracto-pdf", controllers.PdfExtractor)
	}
}
