package model

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	UserID      int
	Amount      float64 `gorm:"type:decimal(10,2)"`
	Currency    string
	PaymentDate time.Time
	Status      string `gorm:"default:'unpaid'"`
}
