package main

import (
	"fmt"
	"geminai_with_go/config"
	"geminai_with_go/models"
	"geminai_with_go/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	ctx_gemini, client_gemini, model, gemini_config, debugResponse, _ := config.GeminiConfig()
	ctx_ollama, client_ollama, _ := config.OllamaConfig()

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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.ApiRoutes(r, ctx_gemini, client_gemini, model, gemini_config, debugResponse, ctx_ollama, client_ollama)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
