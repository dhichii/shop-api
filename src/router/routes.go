package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InitServer() *fiber.App {
	route := fiber.New()

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to Shop API")
	})

	return route
}
