package config

import (
	"fmt"
	"log"
	"os"

	"golang-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB = database

	err = DB.AutoMigrate(
		&domain.User{},
		&domain.ActivityLog{},
	)

	if err != nil {
		log.Fatal("auto migrate failed:", err)
	}

	log.Println("database connected")
}
