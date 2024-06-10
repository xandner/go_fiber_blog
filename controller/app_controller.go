package controller

import (
	"blog/dto"
	"blog/model"
	"blog/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type appController struct {
	userUsecase    usecase.User
	articleUsecase usecase.Article
}

func New(u usecase.User,a usecase.Article) AppController {
	return &appController{u,a}
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
// @Success 200 {string} string "ok"
// @Router /create [post]
func (ac *appController) CreateUser(c *fiber.Ctx) error {
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
	ac.userUsecase.CreateUser(creationData)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// @summary Get Article by ID
// @description Get Article by ID
// @Tags article
// @accept json
// @produce json
// @param id path int true "Article ID"
// @router /article/{id} [get]
func (ac *appController) GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	article, err := ac.articleUsecase.ReadArticleByID(intId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(article)
}
