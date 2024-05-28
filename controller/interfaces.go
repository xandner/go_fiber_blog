package controller

import "github.com/gofiber/fiber/v2"

type (
	AppController interface {
		GetTest(*fiber.Ctx) error
		CreateUser(*fiber.Ctx) error
	}
)