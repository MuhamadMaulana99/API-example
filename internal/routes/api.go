package routes

import (
	"golang-api/internal/handler"
	"golang-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	app.Get(
		"/swagger/*",
		swagger.HandlerDefault,
	)

	api.Post(
		"/register",
		handler.Register,
	)
	// api.Delete(
	// 	"/:id",
	// 	handler.DeleteUser,
	// )

	api.Post(
		"/login",
		handler.Login,
	)

	user := api.Group(
		"/users",
		middleware.JWTProtected(),
	)

	user.Get(
		"/",
		handler.GetUsers,
	)

	user.Put(
		"/:id",
		handler.UpdateUser,
	)

	user.Delete(
		"/:id",
		middleware.AdminOnly(),
		handler.DeleteUser,
	)

	user.Get(
		"/profile",
		handler.Profile,
	)
}
