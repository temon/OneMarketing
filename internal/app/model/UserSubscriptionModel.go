package model

import (
	"gorm.io/gorm"
	"time"
)

type UserSubscription struct {
	gorm.Model
	UserID             int
	SubscriptionPlanID int
	PaymentID          int
	StartDate          time.Time
	EndDate            time.Time
	Status             string
}
