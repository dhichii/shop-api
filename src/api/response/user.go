package response

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	UserParam struct {
		User     *model.User
		Provinsi *Provinsi
		Kota     *Kota
	}

	User struct {
		ID           int      `json:"id"`
		Nama         string   `json:"nama"`
		NoTelp       string   `json:"no_telp"`
		TanggalLahir string   `json:"tanggal_Lahir"`
		Pekerjaan    string   `json:"pekerjaan"`
		JenisKelamin string   `json:"jenis_kelamin"`
		Tentang      string   `json:"tentang"`
		Email        string   `json:"email"`
		UpdatedAt    string   `gorm:"type:date"`
		CreatedAt    string   `gorm:"type:date"`
		Provinsi     Provinsi `json:"id_provinsi"`
		Kota         Kota     `json:"id_kota"`
	}
)

func (p *UserParam) MapToResponse() *User {
	return &User{
		ID:           p.User.ID,
		Nama:         p.User.Nama,
		NoTelp:       p.User.NoTelp,
		TanggalLahir: helper.ConvertDateToString(p.User.TanggalLahir),
		Pekerjaan:    p.User.Pekerjaan,
		JenisKelamin: p.User.JenisKelamin,
		Tentang:      p.User.Tentang,
		Email:        p.User.Email,
		UpdatedAt:    helper.ConvertDateToString(p.User.UpdatedAt),
		CreatedAt:    helper.ConvertDateToString(p.User.CreatedAt),
		Provinsi:     *p.Provinsi,
		Kota:         *p.Kota,
	}
}
