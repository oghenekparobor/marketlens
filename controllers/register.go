package controllers

import (
	"log"
	"net/http"
	"oghenekparobor/market-lens/config"
	"oghenekparobor/market-lens/models"
	"oghenekparobor/market-lens/params"
	"oghenekparobor/market-lens/repositories"
	"oghenekparobor/market-lens/responses"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var rp params.RegisterParam

	if err := ctx.ShouldBindJSON(&rp); err != nil {
		responses.SendResponse(ctx, http.StatusBadRequest, "Invalid input", nil, true)
		return
	}

	// Check password strength
	if err := repositories.CheckPasswordStrength(rp.Password); err != nil {
		responses.SendResponse(ctx, http.StatusBadRequest, err.Error(), nil, true)
		return
	}

	role, err := repositories.FetchRole(config.DB, "customer")
	if err != nil {
		responses.SendResponse(ctx, http.StatusInternalServerError, "Failed to fetch user role", nil, true)
		return
	}

	if exists, err := repositories.DoesEmailExist(config.DB, rp.Email); err != nil {
		responses.SendResponse(ctx, http.StatusInternalServerError, "Error checking email existence", nil, true)
		return
	} else if exists {
		responses.SendResponse(ctx, http.StatusConflict, "Email already exists", nil, true)
		return
	}

	hashedPassword, err := repositories.HashPassword(rp.Password)
	if err != nil {
		responses.SendResponse(ctx, http.StatusBadRequest, "Invalid Password", nil, true)
		return
	}

	user := models.User{
		FirstName:    rp.FirstName,
		LastName:     rp.LastName,
		Email:        rp.Email,
		PasswordHash: hashedPassword,
		PhoneNumber:  rp.PhoneNumber,
		RoleID:       role.ID,
	}

	if err := repositories.CreateUser(config.DB, &user); err != nil {
		log.Printf("Error creating user: %v", err)
		responses.SendResponse(ctx, http.StatusInternalServerError, "Failed to create user", nil, true)
		return
	}

	// todo: auto login the user
	responses.SendResponse(ctx, http.StatusCreated, "User registered successfully!", user, false)
}
