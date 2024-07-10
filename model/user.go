package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	PasswordHash   string
	Bio            *string
	Website        *string
	ProfilePicture *[]byte
	Country        string
	Friends        []*User `gorm:"many2many:user_friends"`
}

func NewUser(name, email, password, country string, bio, website *string, profilePicture *[]byte) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:           name,
		Email:          email,
		PasswordHash:   hashedPassword,
		Bio:            bio,
		Website:        website,
		ProfilePicture: profilePicture,
		Country:        country,
		Friends:        []*User{},
	}, nil
}

func (u *User) ToString() string {
	return fmt.Sprintf("User[ID: %d, Name: %s, Email: %s, Country: %s]", u.ID, u.Name, u.Email, u.Country)
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
