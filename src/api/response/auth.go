package response

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	LoginParam struct {
		User     *model.User
		Provinsi *Provinsi
		Kota     *Kota
		Token    string
	}

	Login struct {
		Nama         string   `json:"nama"`
		NoTep        string   `json:"no_tep"`
		TanggalLahir string   `json:"tanggal_Lahir"`
		Tentang      string   `json:"tentang"`
		Pekerjaan    string   `json:"pekerjaan"`
		Email        string   `json:"email"`
		Provinsi     Provinsi `json:"id_provinsi"`
		Kota         Kota     `json:"id_kota"`
		Token        string   `json:"token"`
	}
)

func (p *LoginParam) MapToResponse() *Login {
	return &Login{
		Nama:         p.User.Nama,
		NoTep:        p.User.NoTelp,
		TanggalLahir: helper.ConvertDateToString(p.User.TanggalLahir),
		Tentang:      p.User.Tentang,
		Pekerjaan:    p.User.Pekerjaan,
		Email:        p.User.Email,
		Provinsi:     *p.Provinsi,
		Kota:         *p.Kota,
		Token:        p.Token,
	}
}
