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

type PostRepository struct {
	DB *gorm.DB
}

func (r *PostRepository) FindById(id string) (*Post, error) {
	var post Post
	if err := r.DB.Find(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Timeline(userIds []string) ([]*Post, error) {
	var posts []*Post
	if err := r.DB.Order("created_at desc").Where("user_id in ?", userIds).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Create(input NewPost, userId string) (string, error) {
	id, _ := uuid.NewRandom()

	newPost := Post{
		ID:        id.String(),
		UserID:    userId,
		Text:      input.Text,
		CreatedAt: time.Now(),
	}
	if err := r.DB.Create(&newPost).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}
