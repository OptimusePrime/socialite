package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserJWTAccessTokenClaims struct {
	UserID uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

type UserJWTRefreshTokenClaims struct {
	UserID uuid.UUID `json:"userId"`
	jwt.StandardClaims
}
