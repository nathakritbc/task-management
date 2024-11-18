package routes

import (
	"task_management_service/config"
	"task_management_service/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(apiV1 fiber.Router) {
	users := apiV1.Group("/users", config.ConfigAuth)

	// Users Route
	users.Get("/:id", handlers.GetUserById)
}
