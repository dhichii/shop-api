package model

import "time"

type (
	Produk struct {
		ID            int    `gorm:"primaryKey"`
		NamaProduk    string `gorm:"type:varchar(255)"`
		Slug          string `gorm:"type:varchar(255)"`
		HargaReseller int
		HargaKonsumen int
		Stok          int
		Deskripsi     string    `gorm:"type:text"`
		CreatedAt     time.Time `gorm:"type:date"`
		UpdatedAt     time.Time `gorm:"type:date"`
		IdToko        int
		IdCategory    int
		Photos        []FotoProduk `gorm:"foreignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}

	FotoProduk struct {
		ID        int `gorm:"primaryKey"`
		IdProduk  int
		Url       string    `gorm:"type:varchar(255)"`
		CreatedAt time.Time `gorm:"type:date"`
		UpdatedAt time.Time `gorm:"type:date"`
	}
)
