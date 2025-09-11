package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string  `gorm:"unique" json:"username"`
	Email      string  `gorm:"unique" json:"email"`
	Password   string  `json:"-"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Address    string  `json:"address"`
	Gender     int64   `json:"gender"`
}
