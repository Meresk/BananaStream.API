package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/livekit/protocol/auth"
	"os"
	"time"
)

var (
	apiKey    string
	apiSecret string
)

func init() {
	envFile := "development.env" // Укажите имя файла по умолчанию

	// Проверьте, передан ли путь к .env через аргументы или переменные окружения
	if customEnv := os.Getenv("ENV_FILE"); customEnv != "" {
		envFile = customEnv
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Warnw("Warning: Failed to load %s file: %v", envFile, err)
	}

	apiKey = os.Getenv("LIVEKIT_API_KEY")
	apiSecret = os.Getenv("LIVEKIT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		log.Fatalf("API_KEY or API_SECRET is missing in environment variables")
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,OPTIONS",
	}))

	app.Post("/getToken", token)

	log.Fatal(app.Listen(":3000"))
}

func token(c *fiber.Ctx) error {
	var request struct {
		Room     string `json:"room"`
		Identity string `json:"identity"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Room == "" || request.Identity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing room or identity",
		})
	}

	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		Room: request.Room,
	}

	switch request.Role {
	case "teacher":
		grant.RoomJoin = true
		grant.RoomCreate = true
		grant.RoomAdmin = true
	case "student":
		grant.RoomJoin = true
		grant.SetCanPublish(false)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role",
		})
	}

	at.SetVideoGrant(grant).SetValidFor(time.Hour * 2).SetIdentity(request.Identity)

	token, err := at.ToJWT()
	if err != nil {
		log.Error("Failed to generate token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
