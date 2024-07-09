package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	Password       string
	Bio            *string
	Website        *string
	ProfilePicture *[]byte
	Country        string
	Friends        []*User `gorm:"many2many:user_friends"`
}

func NewUser(name, email, password, country string, bio, website *string, profilePicture *[]byte) *User {
	return &User{
		Name:           name,
		Email:          email,
		Password:       password,
		Bio:            bio,
		Website:        website,
		ProfilePicture: profilePicture,
		Country:        country,
		Friends:        []*User{},
	}
}
