package models

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}
