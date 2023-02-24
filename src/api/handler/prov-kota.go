package handler

import (
	"net/http"
	"shop-api/src/api/external"
	"shop-api/src/helper"

	"github.com/gofiber/fiber/v2"
)

var method = "GET"

func GetAllProvinsiHandler(c *fiber.Ctx) error {
	data, err := external.GetAllProvinsi()
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

func GetProvinsiByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := external.GetProvinsiByID(id)
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

	if data.ID == "" {
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
			Data:     data,
		},
	)
}

func GetAllKotaByProvinsiIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := external.GetAllKotaByProvinsiID(id)
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

func GetKotaByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := external.GetKotaByID(id)
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

	if data.ID == "" {
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
			Data:     data,
		},
	)
}
