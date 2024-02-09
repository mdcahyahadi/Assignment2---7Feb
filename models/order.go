package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(191)"`
	OrderAt      time.Time
	Items        []Item
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
