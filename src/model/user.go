package model

import "time"

type User struct {
	ID           int       `gorm:"primaryKey"`
	Nama         string    `gorm:"type:varchar(255)"`
	KataSandi    string    `gorm:"type:varchar(255)"`
	NoTelp       string    `gorm:"unique;varchar(255)"`
	TanggalLahir time.Time `gorm:"type:date"`
	JenisKelamin string    `gorm:"type:varchar(255)"`
	Tentang      string    `gorm:"type:text"`
	Pekerjaan    string    `gorm:"type:varchar(255)"`
	Email        string    `gorm:"unique;type:varchar(255)"`
	IdProvinsi   string    `gorm:"type:varchar(255)"`
	IdKota       string    `gorm:"type:varchar(255)"`
	IsAdmin      bool
	UpdatedAt    time.Time `gorm:"type:date"`
	CreatedAt    time.Time `gorm:"type:date"`
	Toko         Toko      `gorm:"foreignKey:IdUser"`
	Alamat       []Alamat  `gorm:"constraint:OnDelete:CASCADE"`
}
