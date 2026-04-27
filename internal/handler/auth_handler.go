package handler

import (
	"golang-api/internal/dto"
	"golang-api/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(
	c *fiber.Ctx,
) error {

	var req dto.RegisterDTO

	if err := c.BodyParser(
		&req,
	); err != nil {

		return c.Status(400).
			JSON(fiber.Map{
				"message": "invalid body",
			})
	}

	user, err :=
		service.Register(
			req,
		)

	if err != nil {
		return c.Status(500).
			JSON(
				fiber.Map{
					"message": err.Error(),
				})
	}

	return c.Status(201).
		JSON(
			fiber.Map{
				"message": "created",
				"data":    user,
			})
}
func DeleteUser(
	c *fiber.Ctx,
) error {

	idParam :=
		c.Params("id")

	id64, _ :=
		strconv.ParseUint(
			idParam,
			10,
			64,
		)

	targetID := uint(id64)

	actorID := uint(1)
	// nanti ambil dari JWT

	err :=
		service.DeleteUser(
			targetID,
			actorID,
		)

	if err != nil {
		return c.Status(500).
			JSON(
				fiber.Map{
					"message": "delete failed",
				})
	}

	return c.JSON(
		fiber.Map{
			"message": "deleted",
		})
}

func Login(
	c *fiber.Ctx,
) error {

	var req dto.LoginDTO

	c.BodyParser(&req)

	token, err :=
		service.Login(req)

	if err != nil {

		return c.Status(401).
			JSON(
				fiber.Map{
					"message": "unauthorized",
				})
	}

	return c.JSON(
		fiber.Map{
			"token": token,
		})
}

func Profile(
	c *fiber.Ctx,
) error {

	return c.JSON(
		fiber.Map{
			"message": "authorized",
		})
}
