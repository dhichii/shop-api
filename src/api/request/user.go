package request

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type UpdateUserRequest struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_Lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

func (r *UpdateUserRequest) MapRequest() *model.User {
	return &model.User{
		Nama:         r.Nama,
		KataSandi:    helper.CreateHash(r.KataSandi),
		NoTelp:       r.NoTelp,
		TanggalLahir: helper.ConvertStringToDate(r.TanggalLahir),
		JenisKelamin: r.JenisKelamin,
		Tentang:      r.Tentang,
		Pekerjaan:    r.Pekerjaan,
		Email:        r.Email,
		IdProvinsi:   r.IdProvinsi,
		IdKota:       r.IdKota,
	}
}
