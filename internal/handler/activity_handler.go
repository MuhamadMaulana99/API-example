package handler

import (
	"strconv"

	"golang-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

// GetActivityLogs godoc
// @Summary Get activity logs
// @Tags Audit
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Actor name"
// @Success 200 {object} map[string]interface{}
// @Router /activity-logs [get]
func GetActivityLogs(
	c *fiber.Ctx,
) error {

	page, _ :=
		strconv.Atoi(
			c.Query("page", "1"),
		)

	limit, _ :=
		strconv.Atoi(
			c.Query("limit", "10"),
		)

	search :=
		c.Query("search", "")

	logs, total, err :=
		service.GetActivityLogs(
			page,
			limit,
			search,
		)

	if err != nil {

		return c.Status(500).JSON(
			fiber.Map{
				"message": "failed",
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"data": logs,
			"meta": fiber.Map{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	)
}
