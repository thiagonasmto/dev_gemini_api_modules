package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Chat struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Conversation []Mensagem `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE" json:"conversation"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
