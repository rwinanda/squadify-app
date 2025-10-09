package models

import (
	"gorm.io/gorm"
)

type UserCategory struct {
	gorm.Model

	UserID uint  `json:"user_id"`
	User   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	CategoryID uint      `json:"category_id"`
	Category   *Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
}
