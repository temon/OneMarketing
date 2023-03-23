package model

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	UserID       uint
	DataSourceID uint
	Name         string
}
