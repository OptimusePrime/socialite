package dto

import "time"

type CreateUserDTO struct {
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// BirthDate holds the value of the "birthDate" field.
	BirthDate time.Time `json:"birthDate,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Biography holds the value of the "biography" field.
	Biography string `json:"biography,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
}
