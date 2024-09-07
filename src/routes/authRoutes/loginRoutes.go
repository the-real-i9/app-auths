package authRoutes

import (
	"appauths/src/handlers/loginHandlers"
	"appauths/src/handlers/loginHandlers/otpLoginHandlers"
	"appauths/src/handlers/loginHandlers/totpLoginHandlers"

	"github.com/gofiber/fiber/v2"
)

// Multi-factor authentication
func MFALogin(router fiber.Router) {
	router.Post("/otp_2fa", otpLoginHandlers.ValidateOTP)

	router.Post("/totp_2fa", totpLoginHandlers.ValidatePasscode)
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
