package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtParser(ctx *fiber.Ctx)(interface{},error){
	user := ctx.Locals("user").(*jwt.Token)
	if user == nil {
		return 0,fiber.NewError(401,"Unauthorized")
	}
	claims:=user.Claims.(jwt.MapClaims)
	// claims["id"]=int(claims["id"].(int32))
	return claims["userID"],nil

}