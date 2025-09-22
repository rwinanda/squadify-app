package repository

import (
	"errors"
	"fmt"
	"squadify-app/config"
	"squadify-app/dto/request"
	"squadify-app/models"

	"gorm.io/gorm"
)

func CreateProfile(input *request.CreateProfileRequest, userID uint) (*models.UserProfile, error) {
	var existingProfile models.UserProfile
	if err := config.DB.Where("user_id = ?", userID).First(&existingProfile).Error; err == nil {
		return nil, fmt.Errorf("profile already exists for this user")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	userProfile := models.UserProfile{
		UserID:     userID,
		FirstName:  input.FirstName,
		MiddleName: input.MiddleName,
		LastName:   input.LastName,
		Latitude:   input.Latitude,
		Longitude:  input.Longitude,
		Address:    input.Address,
		Gender:     input.Gender,
	}

	if err := config.DB.Create(&userProfile).Error; err != nil {
		return nil, err
	}

	return &userProfile, nil
}
