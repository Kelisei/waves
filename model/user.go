package model

import (
	"gorm.io/gorm"
)

// User modelo
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
