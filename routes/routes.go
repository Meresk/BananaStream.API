package routes

import (
	"BananaStream.API/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	//livekit
	app.Post("/getToken", controllers.Token)
	app.Get("/rooms", controllers.Rooms)

	//auth
	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, db)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, db)
	})
}
