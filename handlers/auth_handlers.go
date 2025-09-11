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

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user, err := repository.CreateUser(&input, hashedPassword)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Registration failed")
		return
	}

	data := gin.H{
		"username": user.Username,
		"email":    user.Email,
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
	user, err := repository.GetUserByEmail(input.Email)

	// Check Password
	if err != nil || !utils.CheckPassword(user.Password, input.Password) {
		response.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT
	token, err := utils.CreateJWT(user.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	data := gin.H{
		"token": token,
	}
	response.Success(c, data, "Login successful")
}
