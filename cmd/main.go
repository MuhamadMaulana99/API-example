package main

import (
	"log"

	"golang-api/config"
	"golang-api/internal/errors"
	"golang-api/internal/middleware"
	"golang-api/internal/routes"
	"golang-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(
			".env not loaded",
		)
	}

	logger.InitLogger()

	config.ConnectDB()

	app := fiber.New(
		fiber.Config{
			ErrorHandler: errors.ErrorHandler,
		},
	)

	app.Use(
		middleware.LoggerMiddleware(),
	)

	routes.Setup(app)

	app.Listen(":3000")
}
