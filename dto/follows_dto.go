package dto

import "github.com/google/uuid"

type CreateFollowDTO struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}

type DeleteFollowDTO struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}
