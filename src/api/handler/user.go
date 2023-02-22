package handler

import (
	"net/http"
	"shop-api/src/api/external"
	"shop-api/src/api/request"
	"shop-api/src/api/response"
	"shop-api/src/config/database"
	"shop-api/src/helper"
	"shop-api/src/model"

	"github.com/gofiber/fiber/v2"
)

func GetMyProfileHandler(c *fiber.Ctx) error {
	method := "GET"
	id := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()
	user := &model.User{}

	if err := DB.WithContext(ctx).Omit("kata_sandi").First(user, id).Error; err != nil {
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

	provinsi, err := external.GetProvinsiByID(user.IdProvinsi)
	if err != nil {
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

	kota, err := external.GetKotaByID(user.IdKota)
	if err != nil {
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

	param := response.UserParam{
		User:     user,
		Provinsi: provinsi,
		Kota:     kota,
	}

	return helper.SuccessResponse(
		helper.ResponseParam{
			Ctx:      c,
			HttpCode: http.StatusOK,
			Method:   method,
			Errors:   nil,
			Data:     param.MapToResponse(),
		},
	)
}

func UpdateProfileHandler(c *fiber.Ctx) error {
	method := "UPDATE"
	request := &request.UpdateUser{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	id := helper.GetUserIDFromLocals(c)
	DB := database.InitMySQL()
	ctx := c.Context()

	if err := DB.WithContext(ctx).Where("id", id).First(new(model.User)).Updates(request.MapRequest()).Error; err != nil {
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
