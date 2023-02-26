package model

import "time"

type (
	Trx struct {
		ID               int `gorm:"primaryKey"`
		IdUser           int
		AlamatPengiriman int
		HargaTotal       int
		KodeInvoice      string      `gorm:"type:varchar(255)"`
		MethodBayar      string      `gorm:"type:varchar(255)"`
		UpdatedAt        time.Time   `gorm:"type:date"`
		CreatedAt        time.Time   `gorm:"type:date"`
		Alamat           Alamat      `gorm:"foreignKey:AlamatPengiriman;constraint:OnDelete:SET NULL"`
		DetailTrx        []DetailTrx `gorm:"foreignKey:IdTrx;constraint:OnDelete:CASCADE"`
	}

	DetailTrx struct {
		ID          int `gorm:"primaryKey"`
		IdTrx       int
		IdLogProduk int
		IdToko      int
		Kuantitas   int
		HargaTotal  int
		UpdatedAt   time.Time `gorm:"type:date"`
		CreatedAt   time.Time `gorm:"type:date"`
		LogProduk   LogProduk `gorm:"foreignKey:IdLogProduk;constraint:OnDelete:CASCADE"`
		Toko        Toko      `gorm:"foreignKey:IdToko;constraint:OnDelete:CASCADE"`
	}

	LogProduk struct {
		ID            int `gorm:"primaryKey"`
		IdProduk      int
		NamaProduk    string `gorm:"type:varchar(255)"`
		Slug          string `gorm:"type:varchar(255)"`
		HargaReseller int
		HargaKonsumen int
		Deskripsi     string    `gorm:"type:text"`
		CreatedAt     time.Time `gorm:"type:date"`
		UpdatedAt     time.Time `gorm:"type:date"`
		IdToko        int
		IdCategory    int
		Produk        Produk       `gorm:"foreignKey:IdProduk;constraint:OnDelete:SET NULL"`
		Toko          Toko         `gorm:"foreignKey:IdToko;constraint:OnDelete:SET NULL"`
		Category      Category     `gorm:"foreignKey:IdCategory;constraint:OnDelete:SET NULL"`
		Photos        []FotoProduk `gorm:"foreignKey:IdProduk;references:IdProduk;constraintOnDelete:SET NULL"`
	}
)
