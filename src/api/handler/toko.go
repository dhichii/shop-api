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

func GetMyTokoHandler(c *fiber.Ctx) error {
	method := "GET"
	IdUser := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	data := &response.MyToko{}

	if err := DB.WithContext(ctx).Table("toko").Where("id_user", IdUser).First(data).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusNotFound,
					Method:   method,
					Errors:   []string{"No Data Category"},
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

func UpdateTokoHandler(c *fiber.Ctx) error {
	method := "UPDATE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	request := &request.UpdateToko{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	url, err := helper.ProcessImage(c, file, "toko")
	if err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	idUser := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	condition := fmt.Sprintf("id_user = %d AND id = %d", idUser, id)

	if err := DB.WithContext(ctx).Where(condition).First(new(model.Toko)).Updates(request.MapRequest(url)).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
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

func GetAllTokoHandler(c *fiber.Ctx) error {
	method := "GET"
	DB := database.InitMySQL()
	ctx := c.Context()
	data := []response.Toko{}

	pagination, err := helper.GetPagination(c)
	if err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

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

func GetTokoByIDHandler(c *fiber.Ctx) error {
	method := "GET"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := &response.Toko{}

	if err := DB.WithContext(ctx).First(data, id).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusNotFound,
					Method:   method,
					Errors:   []string{"Toko tidak ditemukan"},
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
