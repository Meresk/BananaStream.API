package routes

import (
	"BananaStream.API/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/getToken", controllers.Token)
	app.Get("/rooms", controllers.Rooms)
}
