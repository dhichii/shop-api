package handler

import (
	"fmt"
	"net/http"
	"shop-api/src/api/request"
	"shop-api/src/api/response"
	"shop-api/src/config/database"
	"shop-api/src/helper"
	"shop-api/src/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllTrxHandler(c *fiber.Ctx) error {
	method := "GET"

	pagination, err := helper.GetPagination(c)
	if err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := []model.Trx{}

	if err := DB.WithContext(ctx).Preload("Alamat").
		Preload("DetailTrx").Preload("DetailTrx.Toko").Preload("DetailTrx.LogProduk").
		Preload("DetailTrx.LogProduk.Toko").Preload("DetailTrx.LogProduk.Category").Preload("DetailTrx.LogProduk.Photos").
		Offset(helper.CountOffset(pagination.Page)).Limit(pagination.Limit).Find(&data).Error; err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusInternalServerError,
				Method:   method,
				Errors:   []string{http.StatusText(http.StatusInternalServerError)},
				Data:     nil,
			},
		)
	}

	pagination.Data = response.MapBatchTrxResponse(data)

	return helper.SuccessResponse(
		helper.ResponseParam{
			Ctx:      c,
			HttpCode: http.StatusOK,
			Method:   method,
			Errors:   nil,
			Data:     pagination,
		},
	)
}

func GetTrxByIDHandler(c *fiber.Ctx) error {
	method := "GET"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := &model.Trx{}

	if err := DB.WithContext(ctx).Preload("Alamat").
		Preload("DetailTrx").Preload("DetailTrx.Toko").Preload("DetailTrx.LogProduk").
		Preload("DetailTrx.LogProduk.Toko").Preload("DetailTrx.LogProduk.Category").Preload("DetailTrx.LogProduk.Photos").
		First(data, id).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusNotFound,
					Method:   method,
					Errors:   []string{"No Data Trx"},
					Data:     nil,
				},
			)
		}

		return helper.FailedResponse(
			helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusInternalServerError,
				Method:   method,
				Errors:   []string{http.StatusText(http.StatusInternalServerError)},
				Data:     nil,
			},
		)
	}

	return helper.SuccessResponse(
		helper.ResponseParam{
			Ctx:      c,
			HttpCode: http.StatusOK,
			Method:   method,
			Errors:   nil,
			Data:     response.MapTrxResponse(data),
		},
	)
}

func CreateTrxHandler(c *fiber.Ctx) error {
	method := "POST"
	request := &request.Trx{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := request.MapRequest()

	trxHargaTotal := 0
	for _, v := range request.DetailTrx {
		hargaTotal := 0
		logProduk := &model.LogProduk{}
		query := fmt.Sprintf(
			`SELECT id as id_produk, nama_produk, slug, harga_reseller,
			harga_konsumen, deskripsi, id_toko, id_category FROM produk 
			WHERE id = %d`, v.IdProduk,
		)
		if err := DB.WithContext(ctx).Raw(query).First(logProduk).Error; err != nil {
			if err.Error() == helper.NOT_FOUND {
				return helper.FailedResponse(
					helper.ResponseParam{
						Ctx:      c,
						HttpCode: http.StatusNotFound,
						Method:   method,
						Errors:   []string{fmt.Sprintf("No Data Product with Id %d", v.IdProduk)},
						Data:     nil,
					},
				)
			}

			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusInternalServerError,
					Method:   method,
					Errors:   []string{http.StatusText(http.StatusInternalServerError)},
					Data:     nil,
				},
			)
		}

		if err := DB.WithContext(ctx).Create(logProduk).Error; err != nil {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusInternalServerError,
					Method:   method,
					Errors:   []string{http.StatusText(http.StatusInternalServerError)},
					Data:     nil,
				},
			)
		}

		hargaTotal += (logProduk.HargaKonsumen * v.Kuantitas)
		data.DetailTrx = append(data.DetailTrx, model.DetailTrx{
			IdLogProduk: logProduk.ID,
			IdToko:      logProduk.IdToko,
			Kuantitas:   v.Kuantitas,
			HargaTotal:  hargaTotal,
		})

		trxHargaTotal += hargaTotal
	}

	idUser := helper.GetUserIDFromLocals(c)
	data.IdUser = idUser
	data.HargaTotal = trxHargaTotal
	data.KodeInvoice = fmt.Sprintf("%d%d", time.Now().UnixMilli(), idUser)

	if err := DB.WithContext(ctx).Create(data).Error; err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusInternalServerError,
				Method:   method,
				Errors:   []string{http.StatusText(http.StatusInternalServerError)},
				Data:     nil,
			},
		)
	}

	return helper.SuccessResponse(
		helper.ResponseParam{
			Ctx:      c,
			HttpCode: http.StatusOK,
			Method:   method,
			Errors:   nil,
			Data:     data.ID,
		},
	)
}
