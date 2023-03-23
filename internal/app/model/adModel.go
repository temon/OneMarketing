package model

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	UserID       uint
	DataSourceID uint
	CampaignID   uint
	Name         string
	AdType       string
	AdFormat     string
}
