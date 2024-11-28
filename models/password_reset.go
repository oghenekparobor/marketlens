package models

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid"` // Foreign key reference
	User      User      `gorm:"foreignKey:UserID"` // One-to-One relationship
	Token     string    `gorm:"unique;not null"` // Unique token for password reset
	ExpiresAt time.Time `json:"expires_at"` // Expiration time for the token
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
