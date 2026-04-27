package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromToken(c *fiber.Ctx) uint {

	user := c.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	return uint(claims["user_id"].(float64))
}
