// @title User API
// @version 1.0
// @description REST API Fiber PostgreSQL

// @host localhost:3000
// @BasePath /api

// ========================
// JWT AUTH FOR SWAGGER
// ========================
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"golang-api/config"
	_ "golang-api/docs"
	"golang-api/internal/errors"
	"golang-api/internal/middleware"
	"golang-api/internal/routes"
	"golang-api/pkg/logger"
	"log"

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
