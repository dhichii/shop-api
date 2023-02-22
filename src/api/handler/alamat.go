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

func GetMyAlamatHandler(c *fiber.Ctx) error {
	method := "GET"
	userID := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	data := []response.Alamat{}

	if err := DB.WithContext(ctx).Find(&data, "user_id", userID).Error; err != nil {
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

func GetAlamatByIDHandler(c *fiber.Ctx) error {
	method := "GET"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	userID := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	data := &response.Alamat{}
	condition := fmt.Sprintf("user_id = %d AND id = %d", userID, id)

	if err := DB.WithContext(ctx).Where(condition).First(data).Error; err != nil {
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

func CreateAlamatHandler(c *fiber.Ctx) error {
	method := "POST"
	request := &request.PostAlamat{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	userID := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	data := request.MapRequest()
	data.UserID = userID

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

func UpdateAlamatHandler(c *fiber.Ctx) error {
	method := "UPDATE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	request := &request.UpdateAlamat{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	userID := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	condition := fmt.Sprintf("user_id = %d AND id = %d", userID, id)

	if err := DB.WithContext(ctx).Where(condition).First(new(model.Alamat)).Updates(request.MapRequest()).Error; err != nil {
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

func DeleteAlamatHandler(c *fiber.Ctx) error {
	method := "DELETE"
	id, err := helper.ConvertID(c.Params("id"))
	if id == 0 {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	userID := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	condition := fmt.Sprintf("user_id = %d AND id = %d", userID, id)

	query := DB.WithContext(ctx).Where(condition).Delete(new(model.Alamat))
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
