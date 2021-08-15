package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NewUser struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	ScreenID   string `json:"screen_id"`
	ScreenName string `json:"screen_name"`
}

type User struct {
	ID                string `json:"id"`
	Email             string `json:"email"`
	ScreenID          string `json:"screen_id"`
	ScreenName        string `json:"screen_name"`
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func CreateUser(db *gorm.DB, input NewUser) (string, error) {
	id, _ := uuid.NewRandom()
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	newUser := User{
		ID:                id.String(),
		Email:             input.Email,
		EncryptedPassword: string(hash),
		ScreenID:          input.ScreenID,
		ScreenName:        input.ScreenName,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}

func FindById(db *gorm.DB, id string) (*User, error) {
	var user User
	if err := db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		ScreenID:   user.ScreenID,
		ScreenName: user.ScreenName,
	}, nil
}

func FindBy(db *gorm.DB, key string, value string) (*User, error) {
	var user User
	if err := db.Find(&user, key+" = ?", value).Error; err != nil {
		return nil, err
	}
	return &User{
		ID:                user.ID,
		Email:             user.Email,
		ScreenID:          user.ScreenID,
		ScreenName:        user.ScreenName,
		EncryptedPassword: user.EncryptedPassword,
	}, nil
}
