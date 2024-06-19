package controller

import (
	"blog/config"
	"blog/dto"
	"blog/model"
	"blog/pkg/utils"
	"blog/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type appController struct {
	userUsecase    usecase.User
	articleUsecase usecase.Article
	cfg *config.Config
}

func New(u usecase.User, a usecase.Article, cfg *config.Config) AppController {
	return &appController{u, a, cfg}
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
// @router /api/v1/app/article/{id} [get]
// @BasicAuth
func (ac *appController) GetArticleByID(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	fmt.Printf("user: %v", claims)
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

func (ac *appController) CreateArticle(c *fiber.Ctx) error{
	userId,err:=utils.JwtParser(c)
	if err!=nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	data := dto.CreteArticleDto{}
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
	user,err:=ac.userUsecase.GetUserByID(userId.(float64))
	if err!=nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	return ac.articleUsecase.CreateArticle(data,user)
}

// @summary Get Articles
// @description Get Articles
// @Tags article
// @accept json
// @produce json
// @param take query int true "take"
// @param skip query int true "skip"
// @router /api/v1/app/article/list-articles [get]
// @Security BearerAuth
// @in header
// @name Authorization
func (ac *appController) GetArticles(c *fiber.Ctx) error{
	fmt.Println("get articles")
	take,err:=strconv.Atoi(c.Query("take"))
	if err!=nil{
		take=10
	}
	skip,err:=strconv.Atoi(c.Query("skip"))
	if err!=nil{
		skip=0
	}
	articles,err:=ac.articleUsecase.ReadArticles(take,skip)
	if err!=nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(articles)
}

// @summary Get User Articles
// @description Get User Articles
// @Tags article
// @accept json
// @produce json
// @param user_id path int true "user_id"
// @param take query int true "take"
// @param skip query int true "skip"
// @router /api/v1/app/article/list-user-articles/{user_id} [get]
// @Security BearerAuth
func (ac *appController) GetUserArticles(c *fiber.Ctx) error{
	userId,err:=strconv.Atoi(c.Params("user_id"))
	if err!=nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	take,err:=strconv.Atoi(c.Query("take"))
	if err!=nil{
		take=10
	}
	skip,err:=strconv.Atoi(c.Query("skip"))
	if err!=nil{
		skip=0
	}
	articles,err:=ac.articleUsecase.ReadUserArticles(take,skip,userId)
	if err!=nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(articles)
}

// @summary Update Article
// @description Update Article
// @Tags article
// @accept json
// @produce json
// @param id path int true "id"
// @router /api/v1/app/article/{id} [put]
// @Security BearerAuth
func (ac *appController) UpdateArticle(c *fiber.Ctx) error{
	userId,err:=utils.JwtParser(c)
	if err!=nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	if userId==nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthorized",
		})
	}
	id,err:=strconv.Atoi(c.Params("id"))
	if err!=nil{
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	data:=model.Article{}
	if err:=c.BodyParser(&data);err!=nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	return ac.articleUsecase.UpdateArticle(id,data)
}