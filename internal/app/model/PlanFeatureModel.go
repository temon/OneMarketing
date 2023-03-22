package model

import "gorm.io/gorm"

type PlanFeature struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Code        string `gorm:"not null; unique"`
	Type        string `gorm:"not null; enum:limit,unlimited"`
}
