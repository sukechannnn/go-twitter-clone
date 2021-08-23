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
	ScreenID   string `json:"screenId"`
	ScreenName string `json:"screenName"`
}

type User struct {
	ID                string `json:"id"`
	Email             string `json:"email"`
	ScreenID          string `json:"screenId"`
	ScreenName        string `json:"screenName"`
	EncryptedPassword string
	CreatedAt         time.Time `json:"CreatedAt"`
	UpdatedAt         time.Time
}

type UserInfo struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	ScreenID   string    `json:"screenId"`
	ScreenName string    `json:"screenName"`
	Following  bool      `json:"Following"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) All(userId string) ([]*UserInfo, error) {
	var allUsers []*UserInfo
	subQuery := r.DB.Select("user_id", "follow_id").Where("user_id = ?", userId).Table("follow_users")
	rows, err := r.DB.Model(&User{}).Select("users.id, users.screen_id, users.screen_name, users.created_at, f.follow_id").Joins("left join (?) as f on f.follow_id = users.id", subQuery).Where("users.id != ?", userId).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id         string
			screenId   string
			screenName string
			createdAt  time.Time
			followId   *string
			following  bool
		)
		err := rows.Scan(&id, &screenId, &screenName, &createdAt, &followId)
		if err != nil {
			return nil, err
		}
		if followId != nil {
			following = true
		} else {
			following = false
		}
		allUsers = append(allUsers, &UserInfo{
			ID:         id,
			ScreenID:   screenId,
			ScreenName: screenName,
			Following:  following,
			CreatedAt:  createdAt,
		})
	}
	return allUsers, nil
}

func (r *UserRepository) Create(input NewUser) (string, error) {
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
	if err := r.DB.Create(&newUser).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}

func (r *UserRepository) FindById(id string) (*User, error) {
	var user User
	if err := r.DB.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		ScreenID:   user.ScreenID,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (r *UserRepository) FindByIds(ids []string) ([]*User, error) {
	var users []*User
	if err := r.DB.Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	var rUsers []*User
	for _, u := range users {
		rUsers = append(rUsers, &User{
			ID:         u.ID,
			Email:      u.Email,
			ScreenID:   u.ScreenID,
			ScreenName: u.ScreenName,
			CreatedAt:  u.CreatedAt,
		})
	}
	return rUsers, nil
}

func (r *UserRepository) FindBy(key string, value string) (*User, error) {
	var user User
	if err := r.DB.Debug().Find(&user, key+" = ?", value).Error; err != nil {
		return nil, err
	}
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		ScreenID:   user.ScreenID,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (r *UserRepository) FindPasswordById(id string) (*User, error) {
	var user User
	if err := r.DB.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &User{
		EncryptedPassword: user.EncryptedPassword,
	}, nil
}
