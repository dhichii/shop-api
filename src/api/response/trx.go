package response

import "shop-api/src/model"

type (
	Trx struct {
		ID          int    `json:"id"`
		HargaTotal  int    `json:"harga_total"`
		KodeInvoice string `json:"kode_invoice"`
		MethodBayar string `json:"method_bayar"`
		Alamat      `json:"alamat_kirim"`
		DetailTrx   []DetailTrx `json:"detail_trx"`
	}

	DetailTrx struct {
		LogProduk  `json:"product"`
		Toko       `json:"toko"`
		Kuantitas  int `json:"kuantitas"`
		HargaTotal int `json:"harga_total"`
	}

	LogProduk struct {
		ID            int    `json:"id"`
		NamaProduk    string `json:"nama_produk"`
		Slug          string `json:"slug"`
		HargaReseller int    `json:"harga_reseller"`
		HargaKonsumen int    `json:"harga_konsumen"`
		Deskripsi     string `json:"deskripsi"`
		Toko          `json:"toko"`
		Category      `json:"category"`
		Foto          []FotoProduk `json:"photos"`
	}
)

func MapBatchTrxResponse(r []model.Trx) (result []Trx) {
	for _, v := range r {
		result = append(result, *MapTrxResponse(&v))
	}

	return
}

func MapTrxResponse(r *model.Trx) *Trx {
	return &Trx{
		ID:          r.ID,
		HargaTotal:  r.HargaTotal,
		KodeInvoice: r.KodeInvoice,
		MethodBayar: r.MethodBayar,
		Alamat: Alamat{
			ID:           r.Alamat.ID,
			JudulAlamat:  r.Alamat.JudulAlamat,
			NamaPenerima: r.Alamat.NamaPenerima,
			NoTelp:       r.Alamat.NoTelp,
			DetailAlamat: r.Alamat.DetailAlamat,
		},
		DetailTrx: MapBatchDetailTrxResponse(r.DetailTrx),
	}
}

func MapBatchDetailTrxResponse(r []model.DetailTrx) (result []DetailTrx) {
	for _, v := range r {
		result = append(result, *MapDetailTrxResponse(v))
	}

	return
}

func MapDetailTrxResponse(r model.DetailTrx) *DetailTrx {
	return &DetailTrx{
		Kuantitas:  r.Kuantitas,
		HargaTotal: r.HargaTotal,
		Toko: Toko{
			ID:       r.Toko.ID,
			NamaToko: r.LogProduk.Toko.NamaToko,
			UrlFoto:  r.LogProduk.Toko.UrlFoto,
		},
		LogProduk: LogProduk{
			ID:            r.LogProduk.ID,
			NamaProduk:    r.LogProduk.NamaProduk,
			Slug:          r.LogProduk.Slug,
			HargaReseller: r.LogProduk.HargaReseller,
			HargaKonsumen: r.LogProduk.HargaKonsumen,
			Deskripsi:     r.LogProduk.Deskripsi,
			Toko: Toko{
				ID:       r.LogProduk.Toko.ID,
				NamaToko: r.LogProduk.Toko.NamaToko,
				UrlFoto:  r.LogProduk.Toko.UrlFoto,
			},
			Category: Category{
				ID:           r.LogProduk.Category.ID,
				NamaCategory: r.LogProduk.Category.NamaCategory,
			},
			Foto: MapBatchFotoResponse(r.LogProduk.Photos),
		},
	}
}

func MapFotoResponse(r model.FotoProduk) *FotoProduk {
	return &FotoProduk{
		ID:       r.ID,
		IdProduk: r.IdProduk,
		Url:      r.Url,
	}
}

func MapBatchFotoResponse(r []model.FotoProduk) (result []FotoProduk) {
	for _, v := range r {
		result = append(result, *MapFotoResponse(v))
	}
	return
}
