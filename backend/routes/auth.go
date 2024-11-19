package routes

import (
	"task_management_service/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(apiV1 fiber.Router) {
	auth := apiV1.Group("/auth")

	// auth route
	// auth.Post("/signIn", handlers.SignIn)
	// auth.Post("/signUp", handlers.SignUp)

	auth.Post("/signIn", handlers.SignIn2Fa)
	auth.Post("/signUp", handlers.SignUp2Fa)

	// เส้นทางสำหรับการเปิดใช้งาน 2FA
	auth.Post("/enable-2fa/:username", handlers.EnableTwoFactor)
}
