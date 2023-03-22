package model

import (
	"gorm.io/gorm"
	"time"
)

type UserSubscriptionHistory struct {
	gorm.Model
	UserID         uint
	SubscriptionID uint
	StartDate      time.Time
	EndDate        time.Time
}
