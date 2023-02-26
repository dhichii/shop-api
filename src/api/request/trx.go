package request

import "shop-api/src/model"

type (
	Trx struct {
		MethodBayar      string      `json:"method_bayar"`
		AlamatPengiriman int         `json:"alamat_kirim"`
		DetailTrx        []DetailTrx `json:"detail_trx"`
	}

	DetailTrx struct {
		IdProduk  int `json:"product_id"`
		Kuantitas int `json:"kuantitas"`
	}
)

func (r *Trx) MapRequest() *model.Trx {
	return &model.Trx{
		MethodBayar:      r.MethodBayar,
		AlamatPengiriman: r.AlamatPengiriman,
	}
}
