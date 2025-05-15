package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rnium/gofiber/handler"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}