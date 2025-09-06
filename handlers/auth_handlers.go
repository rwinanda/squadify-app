package handlers

import (
	"fmt"
	"net/http"
	"squadify-app/config"
	"squadify-app/dto"
	"squadify-app/models"
	"squadify-app/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register user
func Register(c *gin.Context) {
	var input dto.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username:   input.Username,
		Email:      input.Email,
		Password:   string(hashedPassword),
		FirstName:  input.FirstName,
		MiddleName: input.MiddleName,
		LastName:   input.LastName,
		Latitude:   input.Latitude,
		Longitude:  input.Longitude,
		Address:    input.Address,
		Gender:     input.Gender,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Registration successful",
		"username": user.Username,
		"email":    user.Email,
	})
}

// Login User
func Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	var user models.User

	// Find user where "email"
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password are incorrect"})
		return
	}

	// Check Password
	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password are unmatch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfull",
		"user":    user.Password,
	})
}
