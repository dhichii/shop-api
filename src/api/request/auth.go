package request

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	Login struct {
		NoTelp    string `json:"no_telp"`
		KataSandi string `json:"kata_sandi"`
	}

	Register struct {
		Nama         string `json:"nama"`
		KataSandi    string `json:"kata_sandi"`
		NoTelp       string `json:"no_telp"`
		TanggalLahir string `json:"tanggal_lahir"`
		Pekerjaan    string `json:"pekerjaan"`
		Email        string `json:"email"`
		IdProvinsi   string `json:"id_provinsi"`
		IdKota       string `json:"id_kota"`
	}
)

func (r *Register) MapRequest() *model.User {
	return &model.User{
		Nama:         r.Nama,
		KataSandi:    helper.CreateHash(r.KataSandi),
		NoTelp:       r.NoTelp,
		TanggalLahir: helper.ConvertStringToDate(r.TanggalLahir),
		Pekerjaan:    r.Pekerjaan,
		Email:        r.Email,
		IdProvinsi:   r.IdProvinsi,
		IdKota:       r.IdKota,
		Toko: model.Toko{
			NamaToko: r.Nama,
		},
	}
}
