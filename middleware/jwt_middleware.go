package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtMiddleware() fiber.Handler{
	return jwtware.New(jwtware.Config{
		SigningKey: "secret",
	})
}