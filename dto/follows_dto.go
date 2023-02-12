package dto

import "github.com/google/uuid"

type FollowDTO struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}
