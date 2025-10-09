package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name           string         `gorm:"unique" json:"name"`
	UserCategories []UserCategory `gorm:"foreignKey:CategoryID" json:"user_categories"`
}
