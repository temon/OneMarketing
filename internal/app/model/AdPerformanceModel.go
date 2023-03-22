package model

import (
	"gorm.io/gorm"
	"time"
)

type AdPerformance struct {
	gorm.Model
	AdID     uint
	MetricID uint
	Value    float64
	Date     time.Time
}
