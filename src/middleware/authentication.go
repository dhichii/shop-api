package middleware

import (
	"net/http"
	"shop-api/src/helper"

	"github.com/gofiber/fiber/v2"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := helper.ValidateJWT(c)
		if err != nil {
			return helper.FailedResponse(helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusUnauthorized,
				Method:   string(c.Request().Header.Method()),
				Errors:   []string{err.Error()},
				Data:     nil,
			})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
