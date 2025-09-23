package repository

import (
	"errors"
	"fmt"
	"squadify-app/config"
	"squadify-app/dto/request"
	"squadify-app/dto/response"
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

	if err := config.DB.Preload("User").Find(&userProfile).Error; err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func GetProfileByID(userID uint) (*response.GetProfileResponse, error) {
	var existingProfile models.UserProfile

	if err := config.DB.Preload("User").Where("user_id = ?", userID).First(&existingProfile).Error; err != nil {
		return nil, err
	}

	// map fields into response struct
	profileResponse := &response.GetProfileResponse{
		Username:   existingProfile.User.Username,
		Email:      existingProfile.User.Email,
		FirstName:  existingProfile.FirstName,
		MiddleName: existingProfile.MiddleName,
		LastName:   existingProfile.LastName,
		Latitude:   existingProfile.Latitude,
		Longitude:  existingProfile.Longitude,
		Address:    existingProfile.Address,
		Gender:     existingProfile.Gender,
	}

	return profileResponse, nil
}
