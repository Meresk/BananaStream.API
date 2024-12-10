package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

var (
	ApiKey           string
	ApiSecret        string
	DatabaseURL      string
	LivekitServerURL string
	JWTSecret        string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err) // Если файл не найден, программа завершится с ошибкой
	}

	ApiKey = os.Getenv("LIVEKIT_API_KEY")
	ApiSecret = os.Getenv("LIVEKIT_API_SECRET")
	DatabaseURL = os.Getenv("DATABASE_URL")
	LivekitServerURL = os.Getenv("LIVEKIT_SERVER_URL")
	JWTSecret = os.Getenv("JWT_SECRET_KEY")

	if ApiKey == "" || ApiSecret == "" || DatabaseURL == "" || LivekitServerURL == "" || JWTSecret == "" {
		panic("Required environment variables are missing")
	}
}
