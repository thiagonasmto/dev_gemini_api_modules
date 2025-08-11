package routes

import (
	"context"
	"geminai_with_go/controllers"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func UserRoutes(router *gin.Engine, ctx context.Context, client *genai.Client, model string, config *genai.GenerateContentConfig, debugResponse bool) {
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
			controllers.SimpleRequestController(c, ctx, client, model, config, debugResponse)
		})
	}
}
