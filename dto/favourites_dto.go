package dto

import (
	"github.com/google/uuid"
	"socialite/ent"
)

type CreateFavouriteDTO struct {
	PostID uuid.UUID `json:"postId"`
	UserID uuid.UUID `json:"userId"`
}

type FindFavouriteByUserDTO struct {
	UserID uuid.UUID `json:"userId"`
}

type DeleteFavouriteDTO struct {
	UserID uuid.UUID `json:"userId"`
	PostID uuid.UUID `json:"postId"`
}

type FavouriteDTO struct {
	ID   uuid.UUID `json:"id"`
	Post ent.Post  `json:"post"`
	User ent.User  `json:"user"`
}
