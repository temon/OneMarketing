package model

import "gorm.io/gorm"

type SubscriptionPlanFeature struct {
	gorm.Model
	SubscriptionPlanID uint
	PlanFeatureID      uint
	LimitType          string `gorm:"type:enum('day', 'week', 'month', 'year')"`
	LimitAmount        int
}
