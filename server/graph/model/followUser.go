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

func CreateFollowUser(db *gorm.DB, input NewFollowUser, userId string) (string, error) {
	id, _ := uuid.NewRandom()

	newFollowUser := FollowUser{
		ID:        id.String(),
		UserID:    userId,
		FollowID:  input.FollowID,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&newFollowUser).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}

func FindFollowUserById(db *gorm.DB, id string) (*FollowUser, error) {
	var followUser FollowUser
	if err := db.Find(&followUser, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &FollowUser{
		ID:        followUser.ID,
		UserID:    followUser.UserID,
		FollowID:  followUser.FollowID,
		CreatedAt: followUser.CreatedAt,
	}, nil
}
