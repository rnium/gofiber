package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Cannot load .env")
	}
	return os.Getenv(key)
}