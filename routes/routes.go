package routes

import (
	"BananaStream.API/controllers"
	"BananaStream.API/controllers/role"
	"BananaStream.API/controllers/user"
	"BananaStream.API/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	//livekit
	app.Post("/getTeacherToken", middlewares.AuthMiddleware, controllers.TeacherToken)
	app.Post("/getStudentToken", controllers.StudentToken)

	app.Get("/rooms", controllers.Rooms)
	app.Post("/createRoom", middlewares.AuthMiddleware, controllers.CreateRoom)

	//auth
	app.Post("/login", func(c *fiber.Ctx) error {
		return user.Login(c, db)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return user.Register(c, db)
	})
	app.Get("/isAuth", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	//role
	app.Get("/roles", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return role.GetAll(c, db)
	})
	app.Post("/roles", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return role.Create(c, db)
	})
	app.Put("/roles/:id", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return role.Update(c, db)
	})
	app.Delete("/roles/:id", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return role.Delete(c, db)
	})
}
