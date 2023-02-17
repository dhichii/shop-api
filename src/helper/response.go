package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type (
	Response struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Errors  []string    `json:"errors"`
		Data    interface{} `json:"data"`
	}

	ResponseParam struct {
		Ctx      *fiber.Ctx
		HttpCode int
		Method   string
		Errors   []string
		Data     interface{}
	}
)

func SuccessResponse(param ResponseParam) error {
	message := fmt.Sprintf("Success to %s Data", param.Method)
	return param.Ctx.Status(param.HttpCode).
		JSON(
			Response{
				Status:  true,
				Message: message,
				Data:    param.Data,
			},
		)
}

func FailedResponse(param ResponseParam) error {
	message := fmt.Sprintf("Failed to %s Data", param.Method)
	return param.Ctx.Status(param.HttpCode).
		JSON(
			Response{
				Message: message,
				Errors:  param.Errors,
			},
		)
}
