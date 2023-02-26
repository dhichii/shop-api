package model

import "time"

type Alamat struct {
	ID           int `gorm:"primaryKey"`
	IdUser       int
	JudulAlamat  string    `gorm:"type:varchar(255)"`
	NamaPenerima string    `gorm:"type:varchar(255)"`
	NoTelp       string    `gorm:"type:varchar(255)"`
	DetailAlamat string    `gorm:"type:varchar(255)"`
	UpdatedAt    time.Time `gorm:"type:date"`
	CreatedAt    time.Time `gorm:"type:date"`
}
