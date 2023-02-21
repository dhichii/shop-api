package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserIDFromLocals(c *fiber.Ctx) int {
	return int(c.Locals("claims").(jwt.MapClaims)["id"].(float64))
}
