package middleware

import (
	"errors"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rnium/gofiber/config"
)


var Protected = jwtware.New(jwtware.Config{
	SigningKey:   jwtware.SigningKey{Key: []byte(config.Getenv("JWT_SECRET"))},
	ErrorHandler: handleError,
})

func handleError(c *fiber.Ctx, err error) error {
	if errors.Is(err, jwtware.ErrJWTMissingOrMalformed) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"details": "Invalid token",
	})
}
