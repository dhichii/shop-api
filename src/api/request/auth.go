package request

import (
	"shop-api/src/helper"
	"shop-api/src/model"
)

type (
	LoginRequest struct {
		NoTelp    string `json:"no_telp"`
		KataSandi string `json:"kata_sandi"`
	}

	RegisterRequest struct {
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

func MapRegisterRequest(request *RegisterRequest) *model.User {
	return &model.User{
		Nama:         request.Nama,
		KataSandi:    helper.CreateHash(request.KataSandi),
		NoTelp:       request.NoTelp,
		TanggalLahir: helper.ConvertStringToDate(request.TanggalLahir),
		Pekerjaan:    request.Pekerjaan,
		Email:        request.Email,
		IdProvinsi:   request.IdProvinsi,
		IdKota:       request.IdKota,
		Toko: model.Toko{
			NamaToko: request.Nama,
		},
	}
}
