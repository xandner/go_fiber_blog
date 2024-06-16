package routes

import (
	"blog/controller"

	jwtware "github.com/gofiber/contrib/jwt"
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
	articleGroup := app.Group("/article", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			// JWTAlg: jwtware.RS256,
			Key:   []byte("secret"),
		},
	}))
	articleGroup.Get("/:id", c.GetArticleByID)
	articleGroup.Post("/create", c.CreateArticle)
	return app
}
