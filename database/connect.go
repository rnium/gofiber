package database

import (
	"fmt"

	"github.com/rnium/gofiber/config"
	"github.com/rnium/gofiber/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Getenv("DB_HOST"),
		config.Getenv("DB_PORT"),
		config.Getenv("DB_USER"),
		config.Getenv("DB_PASSWORD"),
		config.Getenv("DB_NAME"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("Cannot migrate database")
	}
}