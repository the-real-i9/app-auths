package authRoutes

import (
	"github.com/gofiber/fiber/v2"
)

func OAuth(router fiber.Router) {
	router.Get("/google/auth_url", nil)
	router.Post("/google/login", nil)

	router.Get("/github/auth_url", nil)
	router.Get("/github/login", nil)
}
