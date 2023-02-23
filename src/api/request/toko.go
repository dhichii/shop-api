package request

import "shop-api/src/model"

type UpdateToko struct {
	NamaToko string `form:"nama_toko"`
}

func (r *UpdateToko) MapRequest(url string) *model.Toko {
	return &model.Toko{
		NamaToko: r.NamaToko,
		UrlFoto:  url,
	}
}
