package models

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/datatypes"
)

// Document representa cada chunk salvo no banco
type Document struct {
	ID        string          `gorm:"primaryKey"`
	Content   string          `gorm:"type:text"`
	Metadata  datatypes.JSON  `gorm:"type:jsonb"`
	Embedding pgvector.Vector `gorm:"type:vector(768)"`
}
