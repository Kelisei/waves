package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	Password       string
	ProfilePicture string
	Email          string       `gorm:"unique"`
	LikedArtists   []Artist     `gorm:"many2many:user_liked_artists"`
	LikedAudios    []Audio      `gorm:"many2many:user_liked_audios"`
	LikedGenres    []Genre      `gorm:"many2many:user_liked_genres"`
	Collections    []Collection `gorm:"many2many:user_collections"`
}

func NewUser(username, password, email, profilePath string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:       username,
		Password:       string(hashedPassword),
		Email:          email,
		ProfilePicture: profilePath,
	}

	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
