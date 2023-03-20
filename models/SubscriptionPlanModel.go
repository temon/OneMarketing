package models

import "gorm.io/gorm"

type SubscriptionPlan struct {
	gorm.Model
	Name         string  `gorm:"not null"`
	Description  string  `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	BillingCycle int     `gorm:"not null"`
	TrialPeriod  int     `gorm:"not null"`
	Status       string  `gorm:"not null"`
}

func (s *SubscriptionPlan) TableName() string {
	return "subscription_plans"
}
