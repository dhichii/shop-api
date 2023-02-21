package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserIDFromLocals(c *fiber.Ctx) float64 {
	return c.Locals("claims").(jwt.MapClaims)["id"].(float64)
}
