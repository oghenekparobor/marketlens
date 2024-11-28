package models

import (
	"github.com/google/uuid"
)

type UserRole struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	RoleName    string    `gorm:"unique;not null"`
	Users       []User    `gorm:"foreignKey:RoleID"` // One-to-Many relationship
}
