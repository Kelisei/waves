package model

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Audios      []Audio
	Collections []Collection `gorm:"many2many:artist_collections"`
	Users       []User       `gorm:"many2many:user_liked_artists"`
}

func NewArtist(name string) *Artist {
	return &Artist{Name: name}
}
