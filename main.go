package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rnium/gofiber/database"
	"github.com/rnium/gofiber/router"
)


func main() {
	database.Connect()
	app := fiber.New()
	router.SetupRouters(app)
	log.Fatal(app.Listen(":8080"))
	
}