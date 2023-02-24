package response

type (
	Produk struct {
		ID            int    `json:"id"`
		NamaProduk    string `json:"nama_produk"`
		Slug          string `json:"slug"`
		HargaReseller int    `json:"harga_reseller"`
		HargaKonsumen int    `json:"harga_konsumen"`
		Stok          int    `json:"stok"`
		Deskripsi     string `json:"deskripsi"`
		Toko          `json:"toko"`
		Category      `json:"category"`
		Photos        []FotoProduk `json:"photos"`
	}

	FotoProduk struct {
		ID       int    `json:"id"`
		IdProduk int    `json:"product_id"`
		Url      string `json:"url"`
	}
)
