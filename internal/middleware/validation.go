package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateBody[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var body T

		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid body",
			})
		}

		c.Locals("body", body)

		return c.Next()
	}
}
