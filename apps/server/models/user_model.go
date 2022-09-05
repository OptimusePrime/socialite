package models

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type User struct {
	Model
	Username  string    `gorm:"not null; unique" json:"username"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Name      string    `gorm:"not null" json:"name"`
	Password  string    `gorm:"not null" json:"password"`
	BirthDate time.Time `gorm:"not null" json:"birthDate"`
	Avatar    string    `json:"avatar"`
	Biography string    `json:"biography"`
	Gender    string    `json:"gender"`
}

func GenerateUser() User {
	return User{
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Name:      gofakeit.Name(),
		Password:  gofakeit.Password(true, true, true, true, false, 32),
		BirthDate: gofakeit.Date(),
		Avatar:    gofakeit.Person().Image,
		Biography: gofakeit.LoremIpsumParagraph(3, 5, 12, "\n"),
		Gender:    gofakeit.Gender(),
	}
}
