package middleware

import (
	"fmt"
	"net/http"
	"shop-api/src/config/database"
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

func ProdukAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := string(c.Request().Header.Method())
		idUser := helper.GetUserIDFromLocals(c)
		id, err := helper.ConvertID(c.Params("id"))
		if id == 0 {
			return helper.FailedResponse(
				helper.ResponseParam{Ctx: c, HttpCode: http.StatusNotAcceptable, Method: method, Errors: []string{err.Error()}, Data: nil},
			)
		}

		DB := database.InitMySQL()
		tokoIdUser := 0
		query := fmt.Sprintf(
			`SELECT toko.id_user FROM toko
			JOIN produk on toko.id=produk.id_toko
			WHERE produk.id=%d`,
			id,
		)

		if err := DB.Raw(query).Scan(&tokoIdUser).Error; err != nil {
			if err.Error() == helper.NOT_FOUND {
				return helper.FailedResponse(
					helper.ResponseParam{
						Ctx:      c,
						HttpCode: http.StatusBadRequest,
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

		if idUser != tokoIdUser {
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
