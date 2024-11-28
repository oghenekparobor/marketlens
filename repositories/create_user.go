package repositories

import (
	"errors"
	"fmt"
	"oghenekparobor/market-lens/models"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *models.User) error {
	// Check if the database connection is nil
	if db == nil {
		return errors.New("database connection is nil")
	}

	user.ID = uuid.New() // Generate a new UUID
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func FetchRole(db *gorm.DB, roleName string) (models.UserRole, error) {
	// Check if the database connection is nil
	if db == nil {
		fmt.Printf("DB is nil")
		return models.UserRole{}, errors.New("database connection is nil")
	}

	var userRole models.UserRole

	// Use First to find the role by its name
	err := db.Where("role_name = ?", roleName).First(&userRole).Error
	if err != nil {
		return models.UserRole{}, err // Return an empty UserRole and the error
	}

	return userRole, nil // Return the found userRole and nil error
}

func HashPassword(password string) (string, error) {
	// Generate a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// DoesEmailExist checks if a user with the given email exists in the database.
func DoesEmailExist(db *gorm.DB, email string) (bool, error) {
	var count int64

	err := db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CheckPasswordStrength checks if the password meets the required strength criteria
func CheckPasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Check for at least one uppercase letter
	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one lowercase letter
	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for at least one number
	number := regexp.MustCompile(`[0-9]`)
	if !number.MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	// Check for at least one special character
	specialChar := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",.<>?\\|]`)
	if !specialChar.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
