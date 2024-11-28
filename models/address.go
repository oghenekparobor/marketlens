package models

import (
	"github.com/google/uuid"
)

type Address struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid"` // Foreign key reference
	User      User      `gorm:"foreignKey:UserID"` // One-to-One relationship
	Street    string    `json:"street" gorm:"null"`
	City      string    `json:"city" gorm:"null"`
	State     string    `json:"state" gorm:"null"`
	Country   string    `json:"country" gorm:"null"`
	ZipCode   string    `json:"zip_code" gorm:"null"`
}
