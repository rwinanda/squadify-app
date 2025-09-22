package repository

import (
	"squadify-app/config"
	"squadify-app/dto/request"
	"squadify-app/models"
)

func CreateUser(input *request.RegisterRequest, hashedPassword string) (*models.User, error) {
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(identifier string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
