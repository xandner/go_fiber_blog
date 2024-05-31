package controller

import (
	"blog/model"
	"blog/usecase"
	"net/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type appController struct {
	userUsecase usecase.User
}

func New(u usecase.User) AppController {
	return &appController{u}
}

func (ac *appController) GetTest(c *fiber.Ctx) error {
	if 1 == 2 {
		return fiber.NewError(400, "fff")
	}
	return c.Status(http.StatusOK).SendString("hi")
}

// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "User object that needs to be added"
// @Success 200 {string} string "ok"
// @Router /create [post]
func (ac *appController) CreateUser(c *fiber.Ctx) error {
	data := model.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Printf("data: %v", data)
	ac.userUsecase.CreateUser(data)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
