package routes

import (
	"golang-api/internal/handler"
	"golang-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	api.Post(
		"/register",
		handler.Register,
	)
	api.Delete(
		"/:id",
		handler.DeleteUser,
	)

	api.Post(
		"/login",
		handler.Login,
	)

	user := api.Group(
		"/users",
		middleware.JWTProtected(),
	)

	user.Get(
		"/profile",
		handler.Profile,
	)
}
