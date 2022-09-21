package services

import "errors"

var (
	ErrUsernameNotUnique          = errors.New("user with the specified username already exists")
	ErrEmailNotUnique             = errors.New("user with the specified email already exists")
	ErrFailedHashingPassword      = errors.New("failed hashing a password")
	ErrInvalidPasswordHash        = errors.New("invalid encoded password hash")
	ErrIncompatibleArgon2Version  = errors.New("incompatible argon2 version")
	ErrUnexpectedJWTSigningMethod = errors.New("unexpected JWT signing method")
	ErrInvalidRefreshToken        = errors.New("invalid refresh token")
	ErrInvalidAccessToken         = errors.New("invalid access token")
	ErrInvalidBearerToken         = errors.New("invalid bearer token")
	ErrFailedCreatingUser         = errors.New("failed creating a user")
	ErrInvalidEmail               = errors.New("invalid email address")
	ErrInvalidName                = errors.New("invalid name")
	ErrInvalidUsername            = errors.New("invalid username")
	ErrInvalidPassword            = errors.New("invalid password")
	ErrInvalidGender              = errors.New("invalid gender")
)
