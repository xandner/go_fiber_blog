package controller

import "github.com/gofiber/fiber/v2"

type (
	AppController interface {
		GetTest(*fiber.Ctx) error
		CreateUser(*fiber.Ctx) error
		GetArticleByID (*fiber.Ctx) error	
		// CreateArticle (*fiber.Ctx) error
	}
	AuthController interface {
		Login(*fiber.Ctx) error 
		SignUp(*fiber.Ctx) error
	}
)