package model

import "time"

type FollowUser struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	FollowID  string `json:"followerId"`
	CreatedAt time.Time
}
