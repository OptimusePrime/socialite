package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserDTO struct {
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password  string    `json:"password,omitempty"`
	BirthDate time.Time `json:"birthDate,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Biography string    `json:"biography,omitempty"`
	Gender    string    `json:"gender,omitempty"`
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
