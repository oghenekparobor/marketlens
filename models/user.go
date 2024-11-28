package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName          string          `json:"first_name" gorm:"not null"`
	LastName           string          `json:"last_name" gorm:"not null"`
	Email              string          `json:"email" gorm:"unique;not null"`
	PhoneNumber        *string   `json:"phone_number" gorm:"null"`
	PasswordHash       string          `json:"password_hash" gorm:"not null"`
	IsEmailVerified    bool            `json:"is_email_verified" gorm:"default:false"`
	IsPhoneVerified    bool            `json:"is_phone_verified" gorm:"default:false"`
	CreatedAt          time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	RoleID             uuid.UUID       `gorm:"type:uuid"` // Foreign key reference
	UserRole           UserRole        `gorm:"foreignKey:RoleID"` // One-to-One relationship
	UserSessions       []UserSession    `gorm:"foreignKey:UserID"` // One-to-Many relationship
	Addresses          []Address        `gorm:"foreignKey:UserID"` // One-to-Many relationship
	PasswordResets     []PasswordReset  `gorm:"foreignKey:UserID"` // One-to-Many relationship
	OTPs               []OTP            `gorm:"foreignKey:UserID"` // One-to-Many relationship
}
