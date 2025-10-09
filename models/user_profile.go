package models

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model

	UserID uint  `json:"user_id"`
	User   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Address    string  `json:"address"`

	GenderID uint    `json:"gender_id"`
	Gender   *Gender `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`
}
