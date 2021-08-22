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

type PostInfo struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Text      string    `json:"text"`
	ScreenID  string    `json:"screenId"`
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

func (r *PostRepository) Timeline(userIds []string) ([]*PostInfo, error) {
	var posts []*PostInfo
	r.DB.Order("posts.created_at desc").Model(&Post{}).Select("posts.id, posts.user_id, posts.text, posts.created_at, users.screen_id").Joins("left join users on users.id = posts.user_id").Where("posts.user_id in ?", userIds).Find(&posts)
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
