package model

import (
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
