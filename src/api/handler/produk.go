package handler

import (
	"fmt"
	"net/http"
	"shop-api/src/api/request"
	"shop-api/src/api/response"
	"shop-api/src/config/database"
	"shop-api/src/helper"
	"shop-api/src/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllProdukHandler(c *fiber.Ctx) error {
	method := "GET"

	pagination, err := helper.GetPagination(c)
	if err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := []response.Produk{}

	if err := DB.WithContext(ctx).Offset(helper.CountOffset(pagination.Page)).Limit(pagination.Limit).Find(&data).Error; err != nil {
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

	pagination.Data = data

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

func GetProdukByIDHandler(c *fiber.Ctx) error {
	method := "GET"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := &response.Produk{}

	if err := DB.WithContext(ctx).First(data, id).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusNotFound,
					Method:   method,
					Errors:   []string{"No Data Product"},
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
			Data:     data,
		},
	)
}

func CreateProductHandler(c *fiber.Ctx) error {
	method := "POST"
	request := &request.Produk{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	photos := []model.FotoProduk{}
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["photos"]
		for _, file := range files {
			url, err := helper.ProcessImage(c, file, "produk")
			if err != nil {
				return helper.FailedResponse(
					helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
				)
			}
			photos = (append(photos, model.FotoProduk{Url: url}))
		}
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	idUser := helper.GetUserIDFromLocals(c)
	idToko := 0
	condition := fmt.Sprintf("SELECT id FROM toko WHERE id_user=%d", idUser)

	// get id toko
	if err := DB.WithContext(ctx).Raw(condition).Scan(&idToko); err != nil {
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

	data := request.MapRequest(photos)
	data.IdToko = idToko

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

func UpdateProdukHandler(c *fiber.Ctx) error {
	method := "UPDATE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	request := &request.Produk{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	photos := []model.FotoProduk{}
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["photos"]
		for _, file := range files {
			url, err := helper.ProcessImage(c, file, "produk")
			if err != nil {
				return helper.FailedResponse(
					helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
				)
			}
			photos = (append(photos, model.FotoProduk{Url: url}))
		}
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()

	if err := DB.WithContext(ctx).Where("id", id).Updates(request.MapRequest(photos)).Error; err != nil {
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
			Data:     nil,
		},
	)
}

func DeleteProdukHandler(c *fiber.Ctx) error {
	method := "DELETE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()

	query := DB.WithContext(ctx).Where("id", id).Delete(new(model.Produk))
	if query.Error != nil {
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

	if query.RowsAffected < 1 && query.Error == nil {
		return helper.FailedResponse(
			helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusBadRequest,
				Method:   method,
				Errors:   []string{helper.NOT_FOUND},
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
			Data:     nil,
		},
	)
}
