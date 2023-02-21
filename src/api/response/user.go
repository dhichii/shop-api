package response

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	UserParam struct {
		User     *model.User
		Provinsi *ProvinsiResponse
		Kota     *KotaResponse
	}

	User struct {
		ID           int              `json:"id"`
		Nama         string           `json:"nama"`
		NoTelp       string           `json:"no_telp"`
		TanggalLahir string           `json:"tanggal_Lahir"`
		Pekerjaan    string           `json:"pekerjaan"`
		JenisKelamin string           `json:"jenis_kelamin"`
		Tentang      string           `json:"tentang"`
		Email        string           `json:"email"`
		UpdatedAt    string           `gorm:"type:date"`
		CreatedAt    string           `gorm:"type:date"`
		Provinsi     ProvinsiResponse `json:"id_provinsi"`
		Kota         KotaResponse     `json:"id_kota"`
	}
)

func (up *UserParam) MapToResponse() *User {
	return &User{
		ID:           up.User.ID,
		Nama:         up.User.Nama,
		NoTelp:       up.User.NoTelp,
		TanggalLahir: helper.ConvertDateToString(up.User.TanggalLahir),
		Pekerjaan:    up.User.Pekerjaan,
		JenisKelamin: up.User.JenisKelamin,
		Tentang:      up.User.Tentang,
		Email:        up.User.Email,
		UpdatedAt:    helper.ConvertDateToString(up.User.UpdatedAt),
		CreatedAt:    helper.ConvertDateToString(up.User.CreatedAt),
		Provinsi:     *up.Provinsi,
		Kota:         *up.Kota,
	}
}
