package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rnium/gofiber/handler"
	"github.com/rnium/gofiber/middleware"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)

	// user
	usr_api := api.Group("/user", middleware.Protected)
	usr_api.Get("/me", handler.UserInfo)
}