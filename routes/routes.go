package routes

import (
	"BananaStream.API/controllers"
	"BananaStream.API/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	//livekit
	app.Post("/getToken", controllers.Token)
	app.Get("/rooms", controllers.Rooms)
	app.Post("/createRoom", middlewares.AuthMiddleware, controllers.CreateRoom)

	//auth
	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, db)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, db)
	})
	app.Get("/isAuth", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
