package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name   string `gorm:"unique;not null"`
	Audios []Audio
	Users  []User `gorm:"many2many:user_liked_genres"`
}

func NewGenre(name string) *Genre {
	return &Genre{Name: name}
}
