package model

import (
	"time"

	"gorm.io/gorm"
)

type Audio struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	ArtistID    uint   `gorm:"not null"`
	Artist      Artist
	Album       string
	Duration    int
	Listenings  int
	ReleaseDate time.Time
	GenreID     uint `gorm:"not null"`
	Genre       Genre
	FilePath    string `gorm:"not null"`
	CoverImage  string
	Users       []User       `gorm:"many2many:user_liked_audios"`
	Collections []Collection `gorm:"many2many:audio_collections"`
}

func NewAudio(name string, artistID uint, album string, duration int, listenings int, releaseDate time.Time, genreID uint, filePath string, coverImage string) *Audio {
	return &Audio{
		Name:        name,
		ArtistID:    artistID,
		Album:       album,
		Duration:    duration,
		Listenings:  listenings,
		ReleaseDate: releaseDate,
		GenreID:     genreID,
		FilePath:    filePath,
		CoverImage:  coverImage,
	}
}
