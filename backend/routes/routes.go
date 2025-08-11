package routes

import (
	"geminai_with_go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
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
}
