package model

import "time"

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
