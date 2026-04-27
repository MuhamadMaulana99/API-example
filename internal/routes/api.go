package routes

import (
	"golang-api/internal/dto"
	"golang-api/internal/handler"
	"golang-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	app.Get("/swagger/*", swagger.HandlerDefault)

	api.Post(
		"/register",
		middleware.ValidateBody[dto.RegisterDTO](),
		handler.Register,
	)

	api.Post(
		"/login",
		middleware.ValidateBody[dto.LoginDTO](),
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
		middleware.ValidateBody[dto.UpdateUserDTO](),
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
