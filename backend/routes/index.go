package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// app.Get("/*", configs.ConfigAuth)
	// app.Static("/", "./images")

	// group  route
	apiV1 := app.Group("/api/v1")

	// init routes
	RegisterPostRoutes(apiV1)
	RegisterUserRoutes(apiV1)
	RegisterAuthRoutes(apiV1)

}
