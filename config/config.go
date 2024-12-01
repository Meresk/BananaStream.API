package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

var (
	ApiKey      string
	ApiSecret   string
	DatabaseURL string
)

func init() {
	envFile := "development.env" // Укажите имя файла по умолчанию

	// Проверка, передан ли путь к .env через аргументы или переменные окружения
	if customEnv := os.Getenv("ENV_FILE"); customEnv != "" {
		envFile = customEnv
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Warnw("Warning: Failed to load %s file: %v", envFile, err)
	}

	ApiKey = os.Getenv("LIVEKIT_API_KEY")
	ApiSecret = os.Getenv("LIVEKIT_API_SECRET")
	DatabaseURL = os.Getenv("DATABASE_URL")

	if ApiKey == "" || ApiSecret == "" || DatabaseURL == "" {
		panic("Required environment variables are missing")
	}
}
