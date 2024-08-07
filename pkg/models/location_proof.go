package models

import "time"

type LocationProof struct {
	TxID            uint      `json:"tx_id" gorm:"primaryKey autoIncrement"`
	ID              uint      `json:"id"`
	Address         string    `json:"address"`
	MachineID       uint64    `json:"machine_id"`
	IsRightLocation bool      `json:"is_right_location"`
	IsMembership    bool      `json:"is_membership"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
}
