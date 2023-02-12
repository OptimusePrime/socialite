package dto

import (
	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Username string `json:"username,omitempty" validate:"required,min=3,max=24"`
	Email    string `json:"email,omitempty" validate:"required,email,max=48"`
	Name     string `json:"name,omitempty" validate:"required,min=3,max=24"`
	Password string `json:"password,omitempty" validate:"required,password,min=8,max=16"`
	// BirthDate time.Time `json:"birthDate,omitempty" validate:"required,birthDate"`
	Gender string `json:"gender,omitempty" validate:"required,max=16"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshUserAccessTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}

type UserAccessTokenDTO struct {
	AccessToken string `json:"accessToken"`
}

type CreateUserDocumentDTO struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Biography string    `json:"biography,omitempty"`
}

type UpdateUserDTO struct {
	AccessToken string `json:"accessToken"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Biography   string `json:"biography"`
	Pronouns    string `json:"pronouns"`
	Gender      string `json:"gender"`
}
