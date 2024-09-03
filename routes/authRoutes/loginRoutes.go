package authRoutes

import (
	"appauths/handlers/loginHandlers"

	"github.com/gofiber/fiber/v2"
)

// Multi-factor authentication
func MFALogin(router fiber.Router) {
	router.Post("/otp/email/send_otp", nil)
	router.Post("/otp/email/verify", nil)

	router.Get("/totp/barcode_setupkey", nil)
	router.Post("/totp/verify", nil)
}

// SSO login
func SSO(router fiber.Router) {

}

func Login(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the login process
	router.Post("/cred", loginHandlers.CredLogin)

	router.Route("/mfa", MFALogin)

	router.Route("/sso", SSO)
}
