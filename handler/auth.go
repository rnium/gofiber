package handler

import (
	"fmt"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/rnium/gofiber/database"
	"github.com/rnium/gofiber/model"
)

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("To be implemented")
}


func Register(c *fiber.Ctx) error {
	type Input struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		RePassword string `json:"re_password"`
	}
	inp := Input{}
	if err := c.BodyParser(&inp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": fmt.Sprintf("Input error: %v", err),
		})
	}
	if inp.Password != inp.RePassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": "Passwords doesn't match",
		})
	}
	_, err := mail.ParseAddress(inp.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": "Invalid email address",
		})
	}

	usr := model.User{
		Name: inp.Name,
		Email: inp.Email,
		Password: inp.Password,
	}
	db := database.DB
	if err := db.Create(&usr).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"details": fmt.Sprintf("Cannot create user: %v", err),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}