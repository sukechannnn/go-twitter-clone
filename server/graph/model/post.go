package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewPost struct {
	Text string `json:"text"`
}

func CreatePost(db *gorm.DB, input NewPost, userId string) (string, error) {
	id, _ := uuid.NewRandom()

	newPost := Post{
		ID:        id.String(),
		UserID:    userId,
		Text:      input.Text,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&newPost).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}
