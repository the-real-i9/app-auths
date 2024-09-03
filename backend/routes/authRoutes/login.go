package authRoutes

import (
	"appauths/backend/handlers/loginHandlers"

	"github.com/gofiber/fiber/v2"
)

func Login(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the login process
	router.Post("/user_pass", loginHandlers.Login)

	router.Get("/totp/barcode", nil)
	router.Post("/totp/verify", nil)
}
