package handler

import (
	"fmt"
	"log"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rnium/gofiber/config"
	"github.com/rnium/gofiber/database"
	"github.com/rnium/gofiber/model"
)

func Login(c *fiber.Ctx) error {
	type Input struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	inp := Input{}
	if err := c.BodyParser(&inp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": fmt.Sprintf("Input error: %v", err),
		})
	}
	_, err := mail.ParseAddress(inp.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": "Invalid email address",
		})
	}

	var user model.User
	db := database.DB
	if err := db.First(&user, "email = ?", inp.Email).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"details": "User not found",
		})
	}
	if ok := IsCorrectPassword(user.Password, inp.Password); !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"details": "Cannot login with the credentials",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})
	token_str, err := token.SignedString([]byte(config.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"details": fmt.Sprintf("Cannot sign token. %v", err),
		})
	}
	return c.JSON(fiber.Map{"token": token_str})
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
	hashed_password, err := HashPassword(inp.Password)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"details": fmt.Sprintf("Cannot hash password. %v", err),
		})
	}
	usr := model.User{
		Name: inp.Name,
		Email: inp.Email,
		Password: hashed_password,
	}
	db := database.DB
	if err := db.Create(&usr).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"details": fmt.Sprintf("Cannot create user. %v", err),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func UserInfo(c *fiber.Ctx) error {
	return c.SendStatus(200)
}