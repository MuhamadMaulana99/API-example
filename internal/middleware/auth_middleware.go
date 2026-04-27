package middleware

import (
	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {

	return jwtware.New(
		jwtware.Config{
			SigningKey: []byte(
				"secretkey",
			),
		},
	)
}
