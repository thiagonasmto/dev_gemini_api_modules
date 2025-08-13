package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type MessageDTO struct {
	Content   string    `gorm:"type:text;not null" json:"content"`
	Role      string    `gorm:"type:varchar(100);not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ChatDTO struct {
	Prompt    string           `gorm:"type:text;not null" json:"prompt"`
	History   []ChatMessageDTO `json:"history"`
	CreatedAt time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}

type ChatMessageDTO struct {
	Role  string           `json:"role"`  // ex: "user" ou "model"
	Parts []MessagePartDTO `json:"parts"` // partes da mensagem
}

type MessagePartDTO struct {
	Type string `json:"type"` // "text" ou "image"
	Data string `json:"data"` // texto normal ou base64 da imagem
}

type Message struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ChatID    uuid.UUID `gorm:"type:uuid;not null" json:"chat_id"` // referência para Chat
	Content   string    `gorm:"type:text;not null" json:"content"`
	Sender    string    `gorm:"type:varchar(100);not null" json:"sender"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
