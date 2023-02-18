package handler

import (
	"net/http"
	"shop-api/src/api/external"
	"shop-api/src/api/request"
	"shop-api/src/api/response"
	"shop-api/src/config/database"
	"shop-api/src/helper"
	"shop-api/src/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	method := "POST"
	user := &model.User{}
	request := &request.LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}
	DB := database.InitMySQL()
	ctx := c.Context()

	if err := DB.WithContext(ctx).Find(user, "no_telp", request.NoTelp).Error; err != nil {
		if err.Error() == helper.NOT_FOUND {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusUnauthorized,
					Method:   method,
					Errors:   []string{"No Telp atau kata sandi salah"},
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

	if !helper.ValidateHash(request.KataSandi, user.KataSandi) {
		return helper.FailedResponse(
			helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusUnauthorized,
				Method:   method,
				Errors:   []string{"No Telp atau kata sandi salah"},
				Data:     nil,
			},
		)
	}

	token := helper.GenerateJWT(user.ID, user.IsAdmin)
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

	param := response.LoginParam{
		User:     user,
		Provinsi: provinsi,
		Kota:     kota,
		Token:    token,
	}
	return helper.SuccessResponse(
		helper.ResponseParam{
			Ctx:      c,
			HttpCode: http.StatusOK,
			Method:   method,
			Errors:   nil,
			Data:     param.MapToLoginResponse(),
		},
	)
}

func RegisterHandler(c *fiber.Ctx) error {
	method := "POST"
	user := &request.RegisterRequest{}
	if err := c.BodyParser(user); err != nil {
		return helper.FailedResponse(
			helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
		)
	}

	DB := database.InitMySQL()
	ctx := c.Context()
	if err := DB.WithContext(ctx).Create(request.MapRegisterRequest(user)).Error; err != nil {
		if strings.Contains(err.Error(), "1062") {
			return helper.FailedResponse(
				helper.ResponseParam{
					Ctx:      c,
					HttpCode: http.StatusBadRequest,
					Method:   method,
					Errors:   []string{err.Error()},
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
			Data:     "Register Succeed",
		},
	)
}
