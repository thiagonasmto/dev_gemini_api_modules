package main

import (
	"fmt"
	"geminai_with_go/config"
	"geminai_with_go/models"
	"geminai_with_go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	// ctx, client, model, gemini_config, debugResponse, _ := config.ConfigGemini()

	fmt.Println("Migrando tabelas...")
	if err := config.DB.AutoMigrate(&models.Client{}); err != nil {
		log.Fatalf("erro ao migrar ClientModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Adm{}); err != nil {
		log.Fatalf("erro ao migrar AdmModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Chat{}); err != nil {
		log.Fatalf("erro ao migrar ChatModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Message{}); err != nil {
		log.Fatalf("erro ao migrar MenssagemModel: %v", err)
	}

	r := gin.Default()

	routes.UserRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
