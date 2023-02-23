package router

import (
	"net/http"
	"shop-api/src/api/handler"
	"shop-api/src/middleware"

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

	category := v1.Group("/category", middleware.Authentication())
	category.Get("", handler.GetAllCategoryHandler)
	category.Get("/:id", handler.GetCategoryByIDHandler)
	category.Post("", middleware.GrantAdmin(), handler.CreateCategoryHandler)
	category.Put("/:id", middleware.GrantAdmin(), handler.UpdateCategoryHandler)
	category.Delete("/:id", middleware.GrantAdmin(), handler.DeleteCategoryHandler)

	user := v1.Group("/user", middleware.Authentication())
	user.Get("", handler.GetMyProfileHandler)
	user.Put("", handler.UpdateProfileHandler)

	alamat := v1.Group("/user/alamat").Use(middleware.Authentication())
	alamat.Get("", handler.GetMyAlamatHandler)
	alamat.Get("/:id", handler.GetAlamatByIDHandler)
	alamat.Post("", handler.CreateAlamatHandler)
	alamat.Put("/:id", handler.UpdateAlamatHandler)
	alamat.Delete("/:id", handler.DeleteAlamatHandler)

	toko := v1.Group("/toko").Use(middleware.Authentication())
	toko.Get("/my", handler.GetMyTokoHandler)
	toko.Put("/:id", handler.UpdateTokoHandler)
	toko.Get("", handler.GetAllTokoHandler)
	toko.Get("/:id", handler.GetTokoByIDHandler)

	return app
}
