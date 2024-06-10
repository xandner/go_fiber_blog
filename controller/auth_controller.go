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

// summary: Login
// description: Login
// @tags: auth
// accept: json
// produce: json
// @Router /login [post]
// @Success 200 {string} string "ok"
// @param login body dto.UserLoginDto true "Login Data"
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

// @summary: SignUp
// @description: SignUp
// @tags: auth
// @accept: json
// @produce: json
// @Router /signup [post]
// @Success 200 {string} string "ok"
// @param signup body dto.UserDto true "User Data"
func (a *authController) SignUp(c *fiber.Ctx) error {
	data := dto.UserDto{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
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
		Password: data.Password,
	}
	a.authUsecase.SignUp(creationData)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
