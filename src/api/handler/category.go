package handler

import (
	"net/http"
	"shop-api/src/api/request"
	"shop-api/src/api/response"
	"shop-api/src/config/database"
	"shop-api/src/helper"
	"shop-api/src/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategoryHandler(c *fiber.Ctx) error {
	method := "GET"
	DB := database.InitMySQL()
	ctx := c.Context()
	data := []response.Category{}

	if err := DB.WithContext(ctx).Find(&data).Error; err != nil {
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

func GetCategoryByIDHandler(c *fiber.Ctx) error {
	method := "GET"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := &response.Category{}

	if err := DB.WithContext(ctx).First(data, id).Error; err != nil {
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

func CreateCategoryHandler(c *fiber.Ctx) error {
	method := "POST"
	request := &request.CategoryRequest{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	data := request.MapRequest()

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

func UpdateCategoryHandler(c *fiber.Ctx) error {
	method := "UPDATE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	request := &request.CategoryRequest{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()

	if err := DB.WithContext(ctx).Where("id", id).First(new(model.Category)).Updates(request.MapRequest()).Error; err != nil {
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

func DeleteCategoryHandler(c *fiber.Ctx) error {
	method := "DELETE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()

	query := DB.WithContext(ctx).Where("id", id).Delete(new(model.Category))
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

	if query.RowsAffected <= 0 && query.Error == nil {
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
