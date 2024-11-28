package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid"` // Foreign key reference
	SessionToken string    `gorm:"null"`
	Ip           string    `gorm:"null"`
	UserAgent    string    `gorm:"null"`
	User         User      `gorm:"foreignKey:UserID"` // One-to-One relationship
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	// Additional fields can be added as required
}
