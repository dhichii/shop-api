package response

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
