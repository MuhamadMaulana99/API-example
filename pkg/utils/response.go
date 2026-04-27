package utils

import (
	"golang-api/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(dto.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message string, err interface{}) error {
	return c.Status(status).JSON(dto.ApiResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
