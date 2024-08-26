package model

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	CreatorID   uint
	Creator     User
	Audios      []Audio  `gorm:"many2many:audio_collections"`
	Artists     []Artist `gorm:"many2many:artist_collections"`
}

func NewCollection(name, description string, creatorID uint) *Collection {
	return &Collection{
		Name:        name,
		Description: description,
		CreatorID:   creatorID,
	}
}
