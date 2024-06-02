package controller

import (
	"blog/dto"
	"blog/model"
	"blog/usecase"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authUsecase usecase.Auth
}

func NewAuthController(authUsecase usecase.Auth) AuthController {
	return &authController{authUsecase}
}

func (a *authController) Login(c *fiber.Ctx) error {
	loginData := dto.UserLoginDto{}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(loginData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := a.authUsecase.Login(loginData)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (a *authController) SignUp(c *fiber.Ctx) error {
	data := dto.UserDto{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Printf("data: %v", data)
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		fmt.Printf("ERROR: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	creationData := model.User{
		Name:   data.Name,
		Phone:  data.Phone,
		Family: data.Family,
	}
	a.authUsecase.SignUp(creationData)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
