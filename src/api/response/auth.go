package response

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	LoginParam struct {
		User     *model.User
		Provinsi *ProvinsiResponse
		Kota     *KotaResponse
		Token    string
	}

	LoginResponse struct {
		Nama         string           `json:"nama"`
		NoTelp       string           `json:"no_telp"`
		TanggalLahir string           `json:"tanggal_Lahir"`
		Tentang      string           `json:"tentang"`
		Pekerjaan    string           `json:"pekerjaan"`
		Email        string           `json:"email"`
		Provinsi     ProvinsiResponse `json:"id_provinsi"`
		Kota         KotaResponse     `json:"id_kota"`
		Token        string           `json:"token"`
	}
)

func (lp *LoginParam) MapToLoginResponse() *LoginResponse {
	return &LoginResponse{
		Nama:         lp.User.Nama,
		NoTelp:       lp.User.NoTelp,
		TanggalLahir: helper.ConvertDateToString(lp.User.TanggalLahir),
		Tentang:      lp.User.Tentang,
		Pekerjaan:    lp.User.Pekerjaan,
		Email:        lp.User.Email,
		Provinsi:     *lp.Provinsi,
		Kota:         *lp.Kota,
		Token:        lp.Token,
	}
}
