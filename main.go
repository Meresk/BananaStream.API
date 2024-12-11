package main

import (
	"BananaStream.API/config"
	"BananaStream.API/db/dbConn"
	"BananaStream.API/db/models"
	"BananaStream.API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     "GET,POST,HEAD,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	db := dbConn.Connect()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
