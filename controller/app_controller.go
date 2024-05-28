package controller

import (
	"blog/model"
	"blog/repo"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type appController struct{
	userRepo repo.UserRepo
}

func New() AppController {
	return &appController{}
}

func (ac *appController) GetTest(c *fiber.Ctx) error{
	if 1==2{
		return fiber.NewError(400,"fff")
	}
	return  c.Status(http.StatusOK).SendString("hi")
}

func (ac *appController) CreateUser(c *fiber.Ctx) error {
	data := model.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Printf("data: %v", data)
	if err := ac.userRepo.CreateUser(data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}