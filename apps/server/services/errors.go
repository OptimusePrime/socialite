package services

import "errors"

var (
	ErrUsernameNotUnique         = errors.New("user with the specified username already exists")
	ErrEmailNotUnique            = errors.New("user with the specified email already exists")
	ErrFailedHashingPassword     = errors.New("failed hashing password")
	ErrInvalidPasswordHash       = errors.New("invalid encoded password hash")
	ErrIncompatibleArgon2Version = errors.New("incompatible argon2 version")
)
