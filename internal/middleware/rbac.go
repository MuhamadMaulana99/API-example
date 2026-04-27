package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AdminOnly() fiber.Handler {

	return func(
		c *fiber.Ctx,
	) error {

		user :=
			c.Locals(
				"user",
			).(*jwt.Token)

		claims :=
			user.Claims.(jwt.MapClaims)

		role :=
			claims["role"].(string)

		if role != "admin" {

			return c.Status(403).
				JSON(
					fiber.Map{
						"message": "forbidden",
					})
		}

		return c.Next()
	}
}
