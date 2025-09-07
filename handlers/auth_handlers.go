package handlers

import (
	"fmt"
	"net/http"
	"squadify-app/dto"
	"squadify-app/repository"
	"squadify-app/utils"

	"github.com/gin-gonic/gin"
)

// Register user
func Register(c *gin.Context) {
	var input dto.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := repository.CreateUser(&input, hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"data": gin.H{
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login User
func Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	// Get user by "email"
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check Password
	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfull",
		"data": gin.H{
			"user": user.Email,
		},
	})
}
