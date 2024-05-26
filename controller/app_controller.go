package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type appController struct{}

func New() AppController {
	return &appController{}
}

func (ac *appController) GetTest(c *fiber.Ctx) error{
	if 1==2{
		return fiber.NewError(400,"fff")
	}
	return  c.Status(http.StatusOK).SendString("hi")
}
