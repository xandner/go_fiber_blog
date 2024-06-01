package routes

import (
	"blog/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(c controller.AppController, a controller.AuthController) *fiber.App {
	app := fiber.New()
	app.Get("/", c.GetTest)
	app.Post("/create", c.CreateUser)

	// Auth routes
	app.Post("/login", a.Login)
	app.Post("/signup", a.SignUp)
	return app
}
