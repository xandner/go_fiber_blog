package routes

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(c controller.AppController, a controller.AuthController) *fiber.App {
	app := fiber.New()
	app.Get("/", c.GetTest)
	app.Post("/create", c.CreateUser)

	// Auth routes
	app.Post("/login", a.Login)
	app.Post("/signup", a.SignUp)

	// Article routes
	articleGroup:=app.Group("/article", middleware.JwtMiddleware())
	articleGroup.Get("/:id", c.GetArticleByID)
	return app
}
