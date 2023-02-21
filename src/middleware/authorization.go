package middleware

import (
	"net/http"
	"shop-api/src/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GrantAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("claims")
		user := claims.(jwt.MapClaims)
		if !user["is_admin"].(bool) {
			return helper.FailedResponse(helper.ResponseParam{
				Ctx:      c,
				HttpCode: http.StatusUnauthorized,
				Method:   string(c.Request().Header.Method()),
				Errors:   []string{http.StatusText(http.StatusUnauthorized)},
				Data:     nil,
			})
		}

		return c.Next()
	}
}
