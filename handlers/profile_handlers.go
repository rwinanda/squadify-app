package handlers

import (
	"fmt"
	"net/http"
	"squadify-app/dto/request"
	"squadify-app/dto/response"
	"squadify-app/repository"

	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
	var input request.CreateProfileRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	userIDRaw, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User ID not found")
		return
	}
	userID := userIDRaw.(uint)

	profile, err := repository.CreateProfile(&input, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, fmt.Sprintf("Something went wrong: %v", err))
		return
	}

	response.Success(c, profile, "Login successful")
}
