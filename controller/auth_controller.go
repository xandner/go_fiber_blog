package controller

import (
	"blog/usecase"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authUsecase usecase.Auth
}

func NewAuthController(authUsecase usecase.Auth) AuthController {
	return &authController{authUsecase}
}

func (a *authController) Login(c *fiber.Ctx) error {
	if 1 == 2 {
		return fiber.NewError(400, "error")
	}
	return a.authUsecase.Login()
}

func (a *authController) SignUp(c *fiber.Ctx) error {
	if 1 == 2 {
		return fiber.NewError(400, "error")
	}
	return a.authUsecase.SignUp()
}
