package model

import "time"

type FollowUser struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	FollowerID string `json:"followerId"`
	CreatedAt  time.Time
}
