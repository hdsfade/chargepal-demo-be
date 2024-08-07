package models

import (
	"fmt"
	"time"
)

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Linked     bool   `json:"linked"`
	Email      string `json:"email"`
	UserID     string `json:"user_id"`
	Address    string `json:"address"`
	Membership bool   `json:"membership"`
}

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"password"`
	Linked     bool      `json:"linked"`
	Email      string    `json:"email"`
	UserID     string    `json:"user_id" gorm:"unique"`
	Address    string    `json:"address" gorm:"unique"`
	Membership bool      `json:"membership"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserLink struct {
	Email  string `json:"email" gorm:"primaryKey"`
	UserID string `json:"userId" gorm:"unique"`
}

func GetUserIdByAddress(address string) (uint, error) {
	var user User

	if err := DB.Where("address = ?", address).First(&user).Error; err != nil {
		return 0, fmt.Errorf("get user by address failed: %w", err)
	}

	return user.ID, nil
}
