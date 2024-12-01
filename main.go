package main

import (
	"BananaStream.API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,OPTIONS",
	}))

	//db := dbConn.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
