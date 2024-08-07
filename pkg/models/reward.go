package models

import "time"

type UserReward struct {
	Address   string    `json:"address" gorm:"primaryKey"`
	Reward    string    `json:"reward"`
	Proof     string    `json:"proof"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type NFTReward struct {
	NFT       uint64    `json:"nft" gorm:"primaryKey"`
	Reward    string    `json:"reward"`
	Proof     string    `json:"proof"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
