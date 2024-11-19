package main

import (
	"task_management_service/config"
	"task_management_service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(config.AppConfigFiber)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello api")
	})
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello ,hello word")
	})

	port := config.Config("PORT")

	//config for customization
	app.Use(cors.New(config.CorsConfigDefault))
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed})) // กำหนดระดับของ compression

	config.InitDatabase()
	routes.SetupRoutes(app)

	app.Listen(":" + port)
}
