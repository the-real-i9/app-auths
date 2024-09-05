package authRoutes

import (
	"appauths/handlers/loginHandlers"
	"appauths/handlers/loginHandlers/otpLoginHandlers"
	"appauths/handlers/loginHandlers/totpLoginHandlers"

	"github.com/gofiber/fiber/v2"
)

// Multi-factor authentication
func MFALogin(router fiber.Router) {
	router.Get("/otp_2fa/send_otp", otpLoginHandlers.SendOTP)
	router.Post("/otp_2fa/verify", otpLoginHandlers.VerifyOTP)

	router.Post("/totp_2fa/validate_passcode", totpLoginHandlers.ValidatePasscode)
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
