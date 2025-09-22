package handlers

import (
	"fmt"
	"net/http"
	"squadify-app/dto/request"
	"squadify-app/dto/response"
	"squadify-app/repository"
	"squadify-app/utils"

	"github.com/gin-gonic/gin"
)

// Register user
func Register(c *gin.Context) {
	var input request.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	// Check if email already exists
	existingUserByEmail, _ := repository.GetUserByEmail(input.Email)
	if existingUserByEmail != nil {
		response.Error(c, http.StatusBadRequest, "Email already registered")
		return
	}

	// Check if username already exists
	existingUserByUsername, _ := repository.GetUserByEmail(input.Username)
	if existingUserByUsername != nil {
		response.Error(c, http.StatusBadRequest, "Username already taken")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Create user
	user, err := repository.CreateUser(&input, hashedPassword)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Registration failed")
		return
	}

	// Generate JWT
	token, err := utils.CreateJWT(user.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	data := gin.H{
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	}
	response.Success(c, data, "Registration successful")
}

// Login user
func Login(c *gin.Context) {
	var input request.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	// Get user by email
	user, err := repository.GetUserByEmail(input.Identifier)

	// Check Password
	if err != nil || !utils.CheckPassword(user.Password, input.Password) {
		response.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT
	token, err := utils.CreateJWT(user.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	data := gin.H{
		"token": token,
	}
	response.Success(c, data, "Login successful")
}
