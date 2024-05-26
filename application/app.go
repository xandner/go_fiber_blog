package application

import (
	"blog/controller"
	"blog/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run() {
	// l, zl := logger.New("INFO")
	// go l.Run()

	httpApp := fiber.New()
	httpApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	appController:=controller.New()
	httpApp.Mount("/api/v1/app", routes.AppRoutes(appController))
	httpApp.Listen("127.0.0.1:8080")
}
