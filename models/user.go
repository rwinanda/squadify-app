package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`

	Profiles       []UserProfile  `gorm:"foreignKey:UserID" json:"profiles"`
	UserCategories []UserCategory `gorm:"foreignKey:UserID" json:"user_categories"`
}
