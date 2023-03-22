package model

import (
	"gorm.io/gorm"
	"time"
)

type APIUsage struct {
	gorm.Model
	UserSubscriptionID int
	APIEndpoint        string
	RequestCount       int
	UsageDate          time.Time
}
