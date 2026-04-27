package handler

import (
	"golang-api/internal/dto"
	"golang-api/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary Register user
// @Description create user baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterDTO true "register"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
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

// GetUsers godoc
// @Summary Get all users
// @Description Ambil semua user
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(
	c *fiber.Ctx,
) error {

	users, err :=
		service.GetAllUsers()

	if err != nil {

		return c.Status(500).
			JSON(
				fiber.Map{
					"message": "failed get users",
				})
	}

	return c.JSON(
		fiber.Map{
			"success": true,
			"data":    users,
		})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update data user
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserDTO true "Update Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(
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

	var req dto.UpdateUserDTO

	if err :=
		c.BodyParser(
			&req,
		); err != nil {

		return c.Status(400).
			JSON(
				fiber.Map{
					"message": "invalid body",
				})
	}

	actorID := uint(1)
	// nanti ambil dari JWT

	updated, err :=
		service.UpdateUser(
			uint(id64),
			req.Name,
			req.Email,
			actorID,
		)

	if err != nil {

		return c.Status(500).
			JSON(
				fiber.Map{
					"message": "update failed",
				})
	}

	return c.JSON(
		fiber.Map{
			"message": "updated",
			"data":    updated,
		})
}

// DeleteUser godoc
// @Summary Delete user
// @Security BearerAuth
// @Tags Users
// @Param id path int true "User ID"
// @Success 200
// @Router /users/{id} [delete]
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

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginDTO true "login"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
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
