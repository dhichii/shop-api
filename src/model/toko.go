package model

import "time"

type Toko struct {
	ID        int `gorm:"primaryKey"`
	IdUser    int
	NamaToko  string    `gorm:"type:varchar(255)"`
	UrlFoto   string    `gorm:"type:varchar(255)"`
	UpdatedAt time.Time `gorm:"type:date"`
	CreatedAt time.Time `gorm:"type:date"`
	Produk    []Produk  `gorm:"foreignKey:IdToko;constraint:OnDelete:CASCADE"`
}
