package request

import "shop-api/src/model"

type PostAlamat struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type UpdateAlamat struct {
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

func (r *PostAlamat) MapRequest() *model.Alamat {
	return &model.Alamat{
		JudulAlamat:  r.JudulAlamat,
		NamaPenerima: r.NamaPenerima,
		NoTelp:       r.NoTelp,
		DetailAlamat: r.DetailAlamat,
	}
}

func (r *UpdateAlamat) MapRequest() *model.Alamat {
	return &model.Alamat{
		NamaPenerima: r.NamaPenerima,
		NoTelp:       r.NoTelp,
		DetailAlamat: r.DetailAlamat,
	}
}
