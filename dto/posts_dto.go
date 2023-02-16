package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"socialite/ent"
	"time"
)

type CreatePostDTO struct {
	Caption  string                `json:"caption"`
	Poster   uuid.UUID             `json:"poster"`
	Image    *multipart.FileHeader `json:"image"`
	Location string                `json:"location"`
}

type DeletePostDTO struct {
	ID string `json:"id"`
}

type FindPostsDTO struct {
	Limit int `json:"limit"`
}

type PostDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Caption   string    `json:"caption"`
	Images    []string  `json:"images"`
	Poster    *ent.User `json:"poster"`
	Location  string    `json:"location"`
}
