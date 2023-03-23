package model

import (
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	gorm.Model
	UserID       int
	DataSourceID int
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	Status       string
}
