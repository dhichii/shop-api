package model

import "time"

type Category struct {
	ID           int       `gorm:"primaryKey"`
	NamaCategory string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"type:date"`
	UpdatedAt    time.Time `gorm:"type:date"`
}
