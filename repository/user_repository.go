package repository

import (
	"squadify-app/config"
	"squadify-app/dto"
	"squadify-app/models"
)

func CreateUser(input *dto.RegisterInput, hashedPassword string) (*models.User, error) {
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
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
