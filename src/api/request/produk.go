package request

import (
	"shop-api/src/model"
	"strings"
)

type Produk struct {
	NamaProduk    string `form:"nama_produk"`
	IdCategory    int    `form:"category_id"`
	HargaReseller int    `form:"harga_reseller"`
	HargaKonsumen int    `form:"harga_konsumen"`
	Stok          int    `form:"stok"`
	Deskripsi     string `form:"deskripsi"`
}

func (r *Produk) MapRequest(photos []model.FotoProduk) *model.Produk {
	slug := strings.ToLower(strings.ReplaceAll(r.NamaProduk, " ", "-"))
	return &model.Produk{
		NamaProduk:    r.NamaProduk,
		Slug:          slug,
		IdCategory:    r.IdCategory,
		HargaReseller: r.HargaReseller,
		HargaKonsumen: r.HargaKonsumen,
		Stok:          r.Stok,
		Deskripsi:     r.Deskripsi,
		Photos:        photos,
	}
}
