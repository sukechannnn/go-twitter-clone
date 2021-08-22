package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowUser struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	FollowID  string `json:"followerId"`
	CreatedAt time.Time
}

type NewFollowUser struct {
	FollowID string `json:"followId"`
}

type FollowUserRepository struct {
	DB *gorm.DB
}

func (r *FollowUserRepository) Create(input NewFollowUser, userId string) (string, error) {
	id, _ := uuid.NewRandom()

	newFollowUser := FollowUser{
		ID:        id.String(),
		UserID:    userId,
		FollowID:  input.FollowID,
		CreatedAt: time.Now(),
	}
	if err := r.DB.Create(&newFollowUser).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}

func (r *FollowUserRepository) FindById(id string) (*FollowUser, error) {
	var followUser FollowUser
	if err := r.DB.Find(&followUser, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &FollowUser{
		ID:        followUser.ID,
		UserID:    followUser.UserID,
		FollowID:  followUser.FollowID,
		CreatedAt: followUser.CreatedAt,
	}, nil
}
