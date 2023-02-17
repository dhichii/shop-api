package router

import (
	"net/http"
	"shop-api/src/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitServer() *fiber.App {
	route := fiber.New()
	route.Use(cors.New())

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to Shop API")
	})

	auth := route.Group("/auth")
	auth.Post("/register", handler.RegisterHandler)

	return route
}
