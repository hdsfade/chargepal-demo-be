package models

import "time"

type Asset struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
