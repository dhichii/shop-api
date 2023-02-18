package router

import (
	"net/http"
	"shop-api/src/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitServer() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to Shop API")
	})

	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	auth.Post("/login", handler.LoginHandler)
	auth.Post("/register", handler.RegisterHandler)

	return app
}
