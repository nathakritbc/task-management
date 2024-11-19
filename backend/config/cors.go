package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CorsConfigDefault = cors.Config{
	Next: nil,
	// AllowOrigins:     "*",
	AllowOrigins:     "http://localhost,http://127.0.0.1,http://localhost:3009,http://localhost:4173,http://127.0.0.1:3009,http://127.0.0.1:4173",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           0,
}
