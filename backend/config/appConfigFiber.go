package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var AppConfigFiber = fiber.Config{
	Prefork:      false,
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
}
