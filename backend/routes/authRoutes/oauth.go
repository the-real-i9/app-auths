package authRoutes

import (
	"github.com/gofiber/fiber/v2"
)

func OAuth(router fiber.Router) {
	router.Get("/google/auth_url")
	router.Post("/google/login")

	router.Get("/github/auth_url")
	router.Get("/github/login")
}
