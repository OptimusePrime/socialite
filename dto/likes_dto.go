package dto

import "github.com/google/uuid"

type LikeDTO struct {
	Post uuid.UUID `json:"post"`
	User uuid.UUID `json:"user"`
}
