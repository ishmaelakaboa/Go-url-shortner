package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID        uint `gorm:"primaryKey"`
	Original  string `gorm:"not null"`
	Short     string `gorm:"type:varchar(10);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}