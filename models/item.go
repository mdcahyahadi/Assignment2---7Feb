package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(191)"`
	Description string `gorm:"not null;type:varchar(191)"`
	Quantity    int
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
