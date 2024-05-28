package routes

import (
	"blog/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(c controller.AppController) *fiber.App {
	app := fiber.New()
	app.Get("/", c.GetTest)
	app.Post("/create", c.CreateUser)
	return app
}
