package models

import (
	"math/big"
	"time"
)

type UserAsset struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Charge    string    `json:"charge"`
	Point     string    `json:"point"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TopUpEvent struct {
	TxID      uint      `json:"tx_id" gorm:"primaryKey;autoIncrement"`
	ID        uint      `json:"id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type PointEvent struct {
	TxID      uint      `json:"tx_id" gorm:"primaryKey;autoIncrement"`
	ID        uint      `json:"id"`
	Price     string    `json:"price"`
	Point     string    `json:"point"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type BuyMembershipEvent struct {
	TxID      uint      `json:"tx_id" gorm:"primaryKey;autoIncrement"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func AddCharge(id uint, amount *big.Int) error {
	userAsset := UserAsset{
		ID:     id,
		Charge: "0",
		Point:  "0",
	}

	if err := DB.Where("id = ?", id).First(&userAsset).Error; err != nil {
		userAsset.Charge = amount.String()
		if err = DB.Create(&userAsset).Error; err != nil {
			return err
		}
	} else {
		value := new(big.Int)
		value.SetString(userAsset.Charge, 10)
		value = value.Add(value, amount)
		userAsset.Charge = value.String()
		if err = DB.Updates(&userAsset).Error; err != nil {
			return err
		}
	}

	event := TopUpEvent{
		ID:     id,
		Amount: amount.String(),
	}
	if err := DB.Create(&event).Error; err != nil {
		return err
	}

	return nil
}
