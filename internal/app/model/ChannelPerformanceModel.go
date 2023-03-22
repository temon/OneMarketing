package model

import (
	"gorm.io/gorm"
	"time"
)

type ChannelPerformance struct {
	gorm.Model
	ChannelID uint
	MetricID  uint
	Value     float64
	Date      time.Time
}
