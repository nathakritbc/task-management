package routes

import (
	"task_management_service/config"
	"task_management_service/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPostRoutes(apiV1 fiber.Router) {
	tasks := apiV1.Group("/tasks", config.ConfigAuth)

	// Tasks Route
	tasks.Get("/", handlers.GetAllTasks)
	tasks.Get("/:id", handlers.GetTaskById)
	tasks.Post("/", handlers.CreateTask)
	tasks.Put("/:id", handlers.UpdateTask)
	tasks.Delete("/:id", handlers.DeleteTask)
}
