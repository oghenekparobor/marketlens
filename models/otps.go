package models

import (
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid"` // Foreign key reference
	User      User      `gorm:"foreignKey:UserID"` // One-to-One relationship
	Code      string    `gorm:"not null"` // OTP code
	ExpiresAt time.Time `json:"expires_at"` // Expiration time for the OTP
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
