package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint                `gorm:"primaryKey"`
	UserID          string              `gorm:"size:16" json:"user_id" form:"user_id"`
	U               User                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:Username"`
	Amount          int64               `json:"amount" form:"amount"`
	TransactionDate time.Time           `time_format:"2006-01-02" json:"transaction_date" form:"transaction_date"`
	IsPaid          bool                `gorm:"default:false" json:"is_paid" form:"is_paid"`
	Detail          []TransactionDetail `gorm:"foreignKey:TransactionID; references:ID"`
	CreatedAt       time.Time           `json:"-"`
	UpdatedAt       time.Time           `json:"-"`
	DeletedAt       gorm.DeletedAt      `gorm:"index" json:"-"`
}

type TransactionDetail struct {
	ID            uint   `gorm:"primaryKey"`
	TransactionID string `json:"transaction_id" form:"transaction_id"`
	StuffID       int    `json:"stuff_id" form:"stuff_id"`
	S             Stuff  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StuffID;references:ID"`
	Count         int64  `json:"count" form:"count"`
}

type TransactionSerializer struct {
	UserID string `gorm:"size:16" json:"user_id" form:"user_id"`
	// Amount          int64                         `json:"amount" form:"amount"`
	// TransactionDate time.Time                     `time_format:"2006-01-02" json:"transaction_date" form:"transaction_date"`
	Detail []TransactionDetailSerializer `json:"detail"`
}

type TransactionDetailSerializer struct {
	// TransactionID string `json:"transaction_id" form:"transaction_id"`
	StuffID int   `json:"stuff_id" form:"stuff_id"`
	Count   int64 `json:"count" form:"count"`
}
