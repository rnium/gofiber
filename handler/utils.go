package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rnium/gofiber/database"
	"github.com/rnium/gofiber/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(raw_password string) (string, error) {
	hashed_bytes, err := bcrypt.GenerateFromPassword([]byte(raw_password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed_bytes), err
}

func IsCorrectPassword(hashed_password, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err == nil
}

func getAuthUserId(c *fiber.Ctx) (uid string, ok bool){
	user_token := c.Locals("user").(*jwt.Token)
	claims := user_token.Claims.(jwt.MapClaims)
	uid, ok = claims["user"].(string)
	return uid, ok
}

func getUser(c *fiber.Ctx, uid string) (model.User, error) {
	var user model.User
	db := database.DB
	if err := db.First(&user, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"details": "User not found",
			})
		} else {
			return user, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"details": "Cannot get the user",
			})
		}
	}
	return user, nil
}