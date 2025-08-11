package main

import (
	"fmt"
	"geminai_with_go/config"
	"geminai_with_go/models"
	"log"
)

func main() {
	config.Connect()

	fmt.Println("Migrando tabelas...")
	if err := config.DB.AutoMigrate(&models.Client{}); err != nil {
		log.Fatalf("erro ao migrar UserModel: %v", err)
	}
	// ctx, client, model, config, debugResponse, _ := config.ConfigGemini()
	// ctx, client, model, config, _, _ := config.ConfigGemini()
	// prompt := "Resume this document"
}
