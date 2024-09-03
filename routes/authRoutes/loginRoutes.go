package authRoutes

import (
	"appauths/handlers/loginHandlers"

	"github.com/gofiber/fiber/v2"
)

func CredLogin(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the login process
	router.Post("/", loginHandlers.CredLogin)

}

func MFALogin(router fiber.Router) {
	router.Post("/otp/email")

	router.Get("/totp/barcode", nil)
	router.Post("/totp/verify", nil)
}

func SSO(router fiber.Router) {

}

func Login(router fiber.Router) {
	router.Route("/cred", CredLogin)

	router.Route("/mfa", MFALogin)

	router.Route("/sso", SSO)
}
