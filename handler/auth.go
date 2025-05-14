package handler

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("To be implemented soon")
}